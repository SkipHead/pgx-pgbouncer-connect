package storage

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log/slog"
	"os"
)

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

func (m *Migration) Up() {
	sql, err := goose.OpenDBWithDriver(m.Driver, m.DBString)
	gooseErrors(err, "goose.OpenDBWithDriver")

	err = goose.Up(sql, m.DirString)
	gooseErrors(err, "goose.Up")
}

func (m *Migration) Down() {
	sql, err := goose.OpenDBWithDriver(m.Driver, m.DBString)
	gooseErrors(err, "goose.OpenDBWithDriver")

	err = goose.Down(sql, m.DirString)
	gooseErrors(err, "goose.Down")
}

func (c *Config) Migrate(path string) *Migration {
	var dbString string

	master, err := c.checkRecovery()
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
