package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	schema := `CREATE TABLE place (
        country text,
        city text NULL,
        telcode integer
    );`

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	countryTelcode := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	countryAll := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`

	db.MustExec(countryTelcode, "Hong Kong", 852)
	db.MustExec(countryTelcode, "Singapore", 65)
	db.MustExec(countryAll, "South Africa", "Johannesburg", 27)

	type Place struct {
		Country string
		City    sql.NullString
		Telcode int
	}

	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)

	for rows.Next() {
		var p Place
		err = rows.StructScan(&p)
		enc.Encode(p)

		m := make(map[string]interface{})
		err = rows.MapScan(m)
		enc.Encode(m)
	}

}
