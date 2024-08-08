package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
	"time"
)

// MasterConn - without a read-only transaction check. Return pgx pool connect master.
func (c *Config) MasterConn(ctx context.Context) (*pgxpool.Pool, error) {

	return pgxpool.New(ctx, c.hostSelect("master"))
}

// ReplicaConn - without a read-only transaction check. Return pgx pool connect replica.
func (c *Config) ReplicaConn(ctx context.Context) (*pgxpool.Pool, error) {

	return pgxpool.New(ctx, c.hostSelect("replica"))
}

func (c *Config) checkRecovery(ctx context.Context) (bool, error) {
	var read bool

	master, err := pgxpool.New(ctx, c.hostSelect("master"))
	if err != nil {
		return read, err
	}
	defer master.Close()

	sql := "SELECT pg_is_in_recovery()"
	err = master.QueryRow(ctx, sql).Scan(&read)
	if err != nil {
		return read, err
	}

	return read, nil
}

// GetAutoConn - Deprecate. See method ReliableConn.
func (c *Config) GetAutoConn(ctx context.Context) (*pgxpool.Pool, error) {

	master, err := c.checkRecovery(ctx)
	if err != nil {
		slog.Error(err.Error(), slog.String("checkRecovery", "GetAutoConn"))
	}
	if master {
		return c.MasterConn(ctx)
	}

	return c.ReplicaConn(ctx)
}

// ReliableConn - with a read-only transaction check. Return pgx pool connect master or replica.
func (c *Config) ReliableConn(ctx context.Context) (*pgxpool.Pool, error) {

	master, err := c.checkRecovery(ctx)
	if err != nil {
		slog.Error(err.Error(), slog.String("checkRecovery", "ReliableConn"))
	}
	if master {
		return c.MasterConn(ctx)
	}

	return c.ReplicaConn(ctx)
}

// New - new connect to data base with sql query sample.
func (c *Connection) New() (*Orm, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	defer cancel()

	reliableConn, err := c.StorageConfig.ReliableConn(ctx)
	if err != nil {
		return nil, err
	}

	if c.Timeout == 0 {
		c.Timeout = 15
	}

	return &Orm{
		Table:    fmt.Sprintf("%s.%s", c.Schema, c.TableName),
		KeyField: c.Columns[0],
		Columns:  c.Columns,
		Pool:     reliableConn,
		Timeout:  c.Timeout,
	}, nil
}
