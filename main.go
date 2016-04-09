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
	cols := []string{}
	vals := []interface{}{}

	for col, val := range record {
		cols = append(cols, col)
		vals = append(vals, val)
	}

	sql, args, _ := squirrel.
		Insert(tableName).
		Columns(cols...).
		Values(vals...).
		ToSql()

	log.Println(sql, args)
}

func UpdateRecordsById(tableName string, record Record, ids []string) {

}

func DeleteRecordsById(tableName string, ids []string) {

}

const SCHEMA = `CREATE TABLE book (
	id     TEXT PRIMARY KEY,
	title  TEXT,
	year   NUMBER
);`

type Book struct {
	ID    string
	Title string
	year  int
}

func main() {
	InsertRecord("book", map[string]interface{}{
		"ID":    "1",
		"Title": "El Mal de Montano",
		"Year":  2005,
	})
	// db, err := sqlx.Open("sqlite3", ":memory:")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// _, err = db.Exec(schema)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// userBirthPlace := `INSERT INTO user (name, birth_place) VALUES (?, ?)`
	// userAll := `INSERT INTO user (name, birth_date, birth_place) VALUES (?, ?, ?)`
	//
	// db.MustExec(userBirthPlace, "Mark", "New York")
	// db.MustExec(userBirthPlace, "Karl", "San Francisco")
	// db.MustExec(userAll, "Helen", "1983-03-13 12:23:32.871", "Johannesburg")
	//
	// type User struct {
	// 	Id         int
	// 	Name       string
	// 	BirthDate  null.String `db:"birth_date"`
	// 	BirthPlace null.String `db:"birth_place"`
	// }
	//
	// rows, err := db.Queryx("SELECT * FROM user")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// enc := json.NewEncoder(os.Stdout)
	//
	// for rows.Next() {
	// 	var u User
	// 	err = rows.StructScan(&u)
	// 	enc.Encode(u)
	// }
	//
	// s := []string{}
	// s = append(s, "UPDATE user SET")
	//
	// m := make(map[string]interface{})
	// m["birth_place"] = "London"
	//
	// for key, value := range m {
	// 	s = append(s, fmt.Sprintf("%s = '%+v'", key, value))
	// }
	//
	// s = append(s, "WHERE id IN (?, ?, ?)")
	//
	// db.MustExec(strings.Join(s, " "), 1, 2, 3)
	//
	// rows, err = db.Queryx("SELECT * FROM user")
	// for rows.Next() {
	// 	var u User
	// 	err = rows.StructScan(&u)
	// 	enc.Encode(u)
	// }
}
