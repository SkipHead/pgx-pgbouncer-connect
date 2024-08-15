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

func (c *Config) master() string {

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s %s",
		c.User, c.Password, c.DbName, c.MasterHost, c.MasterPort, c.options())
}

func (c *Config) replica() string {

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s %s",
		c.User, c.Password, c.DbName, c.ReplicaHost, c.ReplicaPort, c.options())
}
