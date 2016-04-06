package main

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, _ := sqlx.Open("sqlite3", ":memory:")
    // db = sqlx.NewDB(sql.Open("sqlite3", ":memory:"), "sqlite3")
    db.Ping()
}
