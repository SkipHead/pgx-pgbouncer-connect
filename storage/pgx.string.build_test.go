package storage

import "testing"

func config() *Config {
	return &Config{
		User:                 "test",
		Password:             "test",
		DbName:               "test",
		SSlMode:              "",
		PoolMaxConns:         "",
		MasterHost:           "master",
		MasterPort:           "6544",
		ReplicaHost:          "replica",
		ReplicaPort:          "6544",
		DefaultQueryExecMode: "",
		Schema:               "test",
	}
}

func TestOptions(t *testing.T) {
	god := config().options()
	wand := "default_query_exec_mode=cache_describe"
	if god != wand {
		t.Errorf("Result was incorrect, got: %s, want: %s.", god, wand)
	}
}

func TestHostSelectReplica(t *testing.T) {
	god := config().hostSelect(config().ReplicaHost)
	wand := "user=test password=test dbname=test default_query_exec_mode=cache_describe host=replica port=6544"
	if god != wand {
		t.Errorf("Result was incorrect, got: %s, want: %s.", god, wand)
	}
}

func TestHostSelectMaster(t *testing.T) {
	god := config().hostSelect(config().MasterHost)
	wand := "user=test password=test dbname=test default_query_exec_mode=cache_describe host=master port=6544"
	if god != wand {
		t.Errorf("Result was incorrect, got: %s, want: %s.", god, wand)
	}
}
