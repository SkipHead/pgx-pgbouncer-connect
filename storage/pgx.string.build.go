package storage

import "fmt"

func (c *Config) options() string {
	var options string

	if c.DefaultQueryExecMode == "" {
		c.DefaultQueryExecMode = "cache_describe"
	}
	options = fmt.Sprintf("default_query_exec_mode=%s", c.DefaultQueryExecMode)

	if c.SSlMode != "" {
		options = fmt.Sprintf("%s ssl_mode=%s", options, c.SSlMode)
	}
	if c.PoolMaxConns != "" {
		options = fmt.Sprintf("%s pool_max_conns=%s", options, c.PoolMaxConns)
	}

	return options
}

func (c *Config) hostSelect(host string) string {
	str := fmt.Sprintf("user=%s password=%s dbname=%s", c.User, c.Password, c.DbName)

	switch host {
	case "master":
		str = fmt.Sprintf("%s host=%s port=%s", str, c.MasterHost, c.MasterPort)
	case "replica":
		str = fmt.Sprintf("%s host=%s port=%s", str, c.ReplicaHost, c.ReplicaPort)
	}

	return str
}
