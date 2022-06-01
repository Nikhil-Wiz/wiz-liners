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
	var conninfo string = "host=localhost port=55000 user=postgres password=postgrespw dbname=wiz sslmode=disable"
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

	result, err := db1.Insert("currencies").Rows(
        goqu.Record{
			"code":"456",
			"name":"fghfg",
            "created_at":time.Now(),
            "modified_at":time.Now(),
        },
    ).Executor().Exec()
    fmt.Println(result,err)
	// val, err := db1.Update("currencies").
	// 	Where(goqu.C("code").Eq("123")).
	// 	Set(goqu.Record{"name": "gfhgf"}).
	// 	Executor().Exec()
	// fmt.Println(val, err)
}
