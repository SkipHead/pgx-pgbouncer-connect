# PGX-PGBOUNCER-CONNECT

For reliable connection of your postgresql cluster.

## Install

``` go get github.com/SkipHead/pgx-pgbouncer-connect ```

## Example Usage

```go
package main

import (
	"context"
	"github.com/SkipHead/pgx-pgbouncer-connect/storage"
	"log"
)

func main() {

	pgb := storage.Config{
		User:        "db_user",
		Password:    "par0l",
		DbName:      "my_db",
		MasterHost:  "192.168.0.1",
		MasterPort:  "6544",
		ReplicaHost: "192.168.0.2r",
		ReplicaPort: "6544",
		Schema:      "db_schema",
	}

	//Migrate table to db
	migrate := pgb.Migrate("./PathToSqlMigrateFiles")
	migrate.Up()

	//Make SQL query
	sql := "SELECT * FROM my_db"
	db, err := pgb.ReliableConn(context.Background())
	if err != nil {
		return
	}
	var value string
	err = db.QueryRow(context.Background(), sql).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
}
```
