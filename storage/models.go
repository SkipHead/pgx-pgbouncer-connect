package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Config struct {
	User                 string `json:"user"`
	Password             string `json:"password"`
	DbName               string `json:"dbname"`
	SSlMode              string `json:"sslmode,omitempty"`
	PoolMaxConns         string `json:"pool_max_conns,omitempty"`
	MasterHost           string `json:"master_host,omitempty"`
	MasterPort           string `json:"master_port,omitempty"`
	ReplicaHost          string `json:"replica_host,omitempty"`
	ReplicaPort          string `json:"replica_port,omitempty"`
	DefaultQueryExecMode string `json:"default_query_exec_mode,omitempty"`
	Schema               string `json:"schema"`
}

type Orm struct {
	Table        string
	KeyField     string
	PageSize     string
	PageIndex    string
	Columns      []string
	Pool         *pgxpool.Pool
	StartDate    time.Time
	EndTime      time.Time
	AfterMinutes int
	Timeout      int
}

type Connection struct {
	Columns       []string
	StorageConfig *Config
	Schema        string
	TableName     string
	Timeout       int
}
