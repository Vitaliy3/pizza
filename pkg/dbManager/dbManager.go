package dbManager

import (
	"log"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

func Init() (err error) {
	conn, err = sqlx.Connect("postgres", "user=postgres password=123 host=localhost port=5432 database=agile sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func Get() *sqlx.DB {
	return conn
}
