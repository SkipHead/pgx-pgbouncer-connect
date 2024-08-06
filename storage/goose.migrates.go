package storage

import (
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

func (m *Migration) Action(action string) {

	sql, err := goose.OpenDBWithDriver(m.Driver, m.DBString)
	gooseErrors(err, "goose.OpenDBWithDriver")

	switch action {
	case "up":
		gooseErrors(goose.Up(sql, m.DirString), "goose.Up")
	case "down":
		gooseErrors(goose.Down(sql, m.DirString), "goose.Up")
	}
}

func (c *Config) Migrate(path string) *Migration {
	var dbString string
	master, err := c.checkRecovery()
	if err != nil {
		gooseErrors(err, "checkRecovery")
	}
	dbString = c.hostSelect("master")

	if master {
		dbString = c.hostSelect("replica")
	}

	return &Migration{
		Driver:    "pgx",
		DBString:  dbString,
		DirString: path,
	}
}
