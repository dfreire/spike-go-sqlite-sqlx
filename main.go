package main

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    schema := `CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer);`

    db, err := sqlx.Open("sqlite3", ":memory:")
    if (err != nil) {
        log.Fatal(err)
    }

    result, err := db.Exec(schema)
    if (err != nil) {
        log.Fatal(err)
    }

    log.Println(result)
}
