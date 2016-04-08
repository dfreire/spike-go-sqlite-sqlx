package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/guregu/null.v3"

	// "github.com/jmoiron/sqlx"
	// "gopkg.in/jmoiron/sqlx.sqlx-v1.1"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	schema := `CREATE TABLE user (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
        name        TEXT,
        birth_date  TEXT NULL,
        birth_place TEXT NULL
    );`

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	userBirthPlace := `INSERT INTO user (name, birth_date) VALUES (?, ?)`
	userAll := `INSERT INTO user (name, birth_date, birth_place) VALUES (?, ?, ?)`

	db.MustExec(userBirthPlace, "Mark", "New York")
	db.MustExec(userBirthPlace, "Karl", "San Francisco")
	db.MustExec(userAll, "Helen", "1983-03-13 12:23:32.871", "Johannesburg")

	type User struct {
		Name       string
		BirthDate  null.String `db:"birth_date"`
		BirthPlace null.String `db:"birth_place"`
	}

	rows, err := db.Queryx("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)

	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		enc.Encode(u)
	}

	s := []string{}
	s = append(s, "UPDATE user SET")

	m := make(map[string]interface{})
	m["birth_place"] = "London"

	for key, value := range m {
		s = append(s, fmt.Sprintf("%s = %+v", key, value))
	}

	s = append(s, "WHERE id in (1, 2, 3)")

	log.Println(strings.Join(s, " "))
}
