package main

import (
    "fmt"

    "github.com/jmoiron/sqlx"
    // _ "github.com/mattn/go-sqlite3"
)

func main() {
    fmt.Println("Hello")

    var db *sqlx.DB

    db = sqlx.Open("sqlite3", ":memory:")
    db = sqlx.NewDB(sql.Open("sqlite3", ":memory:"), "sqlite3")
    err = db.Ping()
}
