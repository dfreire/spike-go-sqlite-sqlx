package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/Masterminds/squirrel.v1"
	"labix.org/v2/mgo/bson"
)

const SCHEMA = `CREATE TABLE IF NOT EXISTS book (
	id     TEXT PRIMARY KEY,
	title  TEXT,
	author TEXT,
	year   NUMBER
);`

type Book struct {
	Id     string      `json:"id",     dbx:"id"`
	Title  string      `json:"title",  dbx:"title"`
	Author null.String `json:"author", dbx:"author"`
	Year   null.Int    `json:"year",   dbx:"year"`
}

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(SCHEMA)

	dbx := sqlx.NewDb(db, "sqlite3")

	id1 := bson.NewObjectId().Hex()
	id2 := bson.NewObjectId().Hex()
	id3 := bson.NewObjectId().Hex()

	insert1 := squirrel.Insert("book").
		SetMap(squirrel.Eq{
			"id":    id1,
			"title": "El Mal de Montano",
			"year":  2002,
		})
	insert1.RunWith(db).Exec()
	fmt.Println(insert1.ToSql())

	insert2 := squirrel.Insert("book").
		SetMap(squirrel.Eq{
			"id":    id2,
			"title": "Doctor Pasavento",
			"year":  2005,
		})
	insert2.RunWith(db).Exec()
	fmt.Println(insert2.ToSql())

	insert3 := squirrel.Insert("book").
		SetMap(squirrel.Eq{
			"id":    id3,
			"title": "Aire de Dylan",
			"year":  2012,
		})
	insert3.RunWith(db).Exec()
	fmt.Println(insert3.ToSql())

	selectAll(dbx)

	update := squirrel.Update("book").
		SetMap(squirrel.Eq{
			"Author": "Enrique Vila-Matas",
		}).
		Where(squirrel.Eq{
			"id": []string{id1, id2, id3},
		})
	update.RunWith(db).Exec()
	fmt.Println(update.ToSql())

	selectAll(dbx)

	delete := squirrel.Delete("book").
		Where(squirrel.Eq{
			"id": []string{id2, id3},
		})
	delete.RunWith(db).Exec()
	fmt.Println(update.ToSql())

	selectAll(dbx)
}

func selectAll(dbx *sqlx.DB) {
	enc := json.NewEncoder(os.Stdout)

	sql, args, _ := squirrel.Select("*").From("book").ToSql()
	rows, _ := dbx.Queryx(sql, args...)
	for rows.Next() {
		var book Book
		rows.StructScan(&book)
		enc.Encode(book)
	}
}
