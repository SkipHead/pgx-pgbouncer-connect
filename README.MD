# PGX-PGBOUNCER-CONNECT

pgx-pgbouncer-connect is an add-on to the pgx driver. 
It allows you to make an automatic choice between two nodes in a PostgreSQL cluster, 
which is always necessary in high availability systems.
It also combines database migration using the goose tool, 
doing everything in a single package. It is very comfortable.
ORM functionality is added as needed, but raw queries are often used.

## Install

``` go get github.com/SkipHead/pgx-pgbouncer-connect ```

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/SkipHead/pgx-pgbouncer-connect/storage"
	"log"
)

func main() {

	config := storage.Config{
		User:        "db_user",
		Password:    "par0l",
		DbName:      "my_db",
		MasterHost:  "192.168.0.1",
		MasterPort:  "6544",
		ReplicaHost: "192.168.0.2",
		ReplicaPort: "6544",
		Schema:      "db_schema",
	}

	orm := storage.Connection{
		Columns:       []string{"field1", "field2", "field3"},
		StorageConfig: &config,
		Schema:        "my_schema",
		TableName:     "table",
		Timeout:       15,
	}

	//Migrate table to db
	migrate := config.Migrate("./PathToSqlMigrateFiles")
	migrate.Up()



	//Make SQL query
	sql := "SELECT * FROM my_db"
	var value string
	//Raw query 
	pool, _ := config.ReliableConn(context.Background())
	err := pool.QueryRow(context.Background(), sql).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	//Result select
	fmt.Println(value)

	//ORM query INSERT data to table
	db, err := orm.New()
	q, err := db.Pool.Exec(context.Background(), db.Insert(), "1", "2")

	//Result true or false
	fmt.Println(q.Insert())
}
```
