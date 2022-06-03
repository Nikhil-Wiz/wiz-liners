package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	pgdb "wiz-liners/internal/repositories/pgdb"
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

	//run each functions seperately to complete crud operations 
	//later test it on other tables

	val, _:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).Insert("40","pqrst")
	fmt.Println(val)

	// val3, _,_:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).ReadOne("40")
	// fmt.Println(val3.Code,val3.Name,val3.Created_At,val3.Modified_At)


	// var val1 uint = 1
	// val2,_ :=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).ReadMany(&val1,5)
	// for _,val := range val2 {
	// 	fmt.Println(val.Code,val.Name)
	// }

	// val6 := "xyzw"
	// val5, _:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).Update("40",&val6)
	// fmt.Println(val5)

	// err1 :=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).Delete("40")
	// fmt.Println(err1)

}
