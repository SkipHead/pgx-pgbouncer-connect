package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
	"time"
)

func (c *Config) Conn(ctx context.Context, host string) (*pgxpool.Pool, error) {
	if host == "" {
		host = "replica"
	}

	return pgxpool.New(ctx, c.hostSelect(host))
}

func (c *Config) checkRecovery() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var read bool

	master, err := c.Conn(ctx, "master")
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
	master, err := c.checkRecovery()
	if err != nil {
		slog.Error(err.Error(), slog.String("checkRecovery", "GetAutoConn"))
	}
	if master {
		return c.Conn(ctx, "master")
	}

	return c.Conn(ctx, "")
}

// ReliableConn - with a read-only transaction check. Return pgx pool connect master or replica.
func (c *Config) ReliableConn(ctx context.Context) (*pgxpool.Pool, error) {
	master, err := c.checkRecovery()
	if err != nil {
		slog.Error(err.Error(), slog.String("checkRecovery", "GetAutoConn"))
	}
	if master {
		return c.Conn(ctx, "master")
	}

	return c.Conn(ctx, "")
}

// MasterConn - without a read-only transaction check. Return pgx pool connect master.
func (c *Config) MasterConn(ctx context.Context) (*pgxpool.Pool, error) {

	return c.Conn(ctx, "master")
}

// ReplicaConn - without a read-only transaction check. Return pgx pool connect replica.
func (c *Config) ReplicaConn(ctx context.Context) (*pgxpool.Pool, error) {

	return c.Conn(ctx, "")
}
