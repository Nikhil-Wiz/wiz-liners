package main

import (
	"database/sql"
	"embed"
	_"fmt"
	_"wiz-liners/internal/repositories/pgdb"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	var conninfo string = "host=localhost port=55001 user=postgres password=postgrespw dbname=wiz sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	// setup database
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	// run app

	//fmt.Println((&pgdb.PortsRepository{}).Insert("nik", "nikhil", "sea", 123, "tn", 87.76, 78.56))
}
