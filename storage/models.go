package storage

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
