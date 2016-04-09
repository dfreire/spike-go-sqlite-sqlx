package main

import (
	"log"

	"gopkg.in/Masterminds/squirrel.v1"
)

// "github.com/jmoiron/sqlx"
// "gopkg.in/jmoiron/sqlx.sqlx-v1.1"

// _ "github.com/mattn/go-sqlite3"

type Record map[string]interface{}

func InsertRecord(tableName string, record Record) {
	sql, args, _ := squirrel.Insert(tableName).SetMap(record).ToSql()
	log.Println(sql, args)
}

func UpdateRecordsById(tableName string, record Record, ids []string) {
	sql, args, _ := squirrel.Update(tableName).Table(tableName).SetMap(record).Where(squirrel.Eq{"id": ids}).ToSql()
	log.Println(sql, args)
}

func DeleteRecordsById(tableName string, ids []string) {

}

const SCHEMA = `CREATE TABLE book (
	id     TEXT PRIMARY KEY,
	title  TEXT,
	author TEXT,
	year   NUMBER
);`

type Book struct {
	ID     string
	Title  string
	Author string
	year   int
}

func main() {
	InsertRecord("book", Record{
		"ID":    "1",
		"Title": "El Mal de Montano",
		"Year":  2002,
	})
	InsertRecord("book", Record{
		"ID":    "2",
		"Title": "Doctor Pasavento",
		"Year":  2005,
	})
	UpdateRecordsById("book", Record{
		"Author": "Enrique Vila-Matas",
	}, []string{"1", "2"})

}
