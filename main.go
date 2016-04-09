package main

import (
	"log"

	"labix.org/v2/mgo/bson"

	"gopkg.in/Masterminds/squirrel.v1"
)

// "github.com/jmoiron/sqlx"
// "gopkg.in/jmoiron/sqlx.sqlx-v1.1"

// _ "github.com/mattn/go-sqlite3"

const SCHEMA = `CREATE TABLE book (
	id     TEXT PRIMARY KEY,
	title  TEXT,
	author TEXT,
	year   NUMBER
);`

func main() {
	id1 := bson.NewObjectId().Hex()
	id2 := bson.NewObjectId().Hex()

	sql, args, _ := squirrel.
		Insert("book").
		SetMap(squirrel.Eq{
			"id":    id1,
			"title": "El Mal de Montano",
			"year":  2002,
		}).
		ToSql()
	log.Println(sql, args)

	sql, args, _ = squirrel.
		Insert("book").
		SetMap(squirrel.Eq{
			"id":    id2,
			"title": "Doctor Pasavento",
			"year":  2005,
		}).
		ToSql()
	log.Println(sql, args)

	sql, args, _ = squirrel.
		Update("book").
		SetMap(squirrel.Eq{
			"Author": "Enrique Vila-Matas",
		}).
		Where(squirrel.Eq{"id": []string{id1, id2}}).
		ToSql()
	log.Println(sql, args)

	sql, args, _ = squirrel.
		Delete("book").
		Where(squirrel.Eq{"id": []string{id1, id2}}).
		ToSql()
	log.Println(sql, args)
}
