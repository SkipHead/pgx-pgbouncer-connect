# PGX-PGBOUNCER-CONNECT

For reliable connection of your postgresql cluster.

## Install

```go get github.com/SkipHead/pgx-pgbouncer-connect```

## How to use 

``` package main ```

```import "github.com/SkipHead/pgx-pgbouncer-connect/storage"```

```func main() {```
``` ```
```    	conn := storage.Config{```
```		User: "db_user",```
```		Password: "par0l",```
```		DbName: "my_db",```
```		MasterHost: "192.168.0.1",```
```		MasterPort: "6544",```
```		ReplicaHost: "192.168.0.2",```
```		ReplicaPort: "6544",```
```		Schema: "db_schema",```
```	}```
```	```
```	```
```	db, err := conn.ReliableConn(context.Background())```
```	if err != nil {```
```		return```
```	}```
```	var value string```
```	err := db.QueryRow(context.Background(), "SELECT * FROM my_db").Scan(&value)```
```	if err != nil {```
```		log.Fatal(err)```
```	}```
```	log.Println(value)```
```
```} ```