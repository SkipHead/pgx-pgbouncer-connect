package storage

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log/slog"
	"os"
	"time"
)

// Migration config for  database  new migration.
type Migration struct {
	Driver    string
	DBString  string
	DirString string
}

func gooseErrors(e error, value string) {
	if e != nil {
		slog.Error(e.Error(), slog.String("action", value))
		os.Exit(1)
	}
}

// Up creating or modifying a table in the database.
func (m *Migration) Up() {
	sql, err := goose.OpenDBWithDriver(m.Driver, m.DBString)
	gooseErrors(err, "goose.OpenDBWithDriver")

	err = goose.Up(sql, m.DirString)
	gooseErrors(err, "goose.Up")
}

// Down performing drop operations.
func (m *Migration) Down() {
	sql, err := goose.OpenDBWithDriver(m.Driver, m.DBString)
	gooseErrors(err, "goose.OpenDBWithDriver")

	err = goose.Down(sql, m.DirString)
	gooseErrors(err, "goose.Down")
}

// Migrate creates a new configuration for database migration.
func (c *Config) Migrate(path string) *Migration {
	var dbString string

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	master, err := c.checkRecovery(ctx)
	if err != nil {
		gooseErrors(err, "checkRecovery")
	}

	switch master {
	case true:
		dbString = c.hostSelect("replica")
	case false:
		dbString = c.hostSelect("master")
	}

	return &Migration{
		Driver:    "pgx",
		DBString:  dbString,
		DirString: path,
	}
}
