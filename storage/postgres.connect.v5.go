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
