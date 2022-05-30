package main

import (
	"database/sql"
	"embed"
	"fmt"
	"time"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	dialect := goqu.Dialect("postgres")
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


	db1 := dialect.DB(db)
    result, err := db1.Insert("countries").Rows(
        goqu.Record{
			"name":"USA",
			"iso_code":"123",
			"currency_code":"456",
            "created_at":time.Now(),
            "modified_at":time.Now(),
        },
    ).Executor().Exec()
    fmt.Println(result,err)

	
	// run app

	//fmt.Println((&pgdb.PortsRepository{}).Insert("nik", "nikhil", "sea", 123, "tn", 87.76, 78.56))
}
