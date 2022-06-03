package countires

import (
	"database/sql"
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb"

	"github.com/doug-martin/goqu/v9"
)

func AddValue(code , name string )(string , error){

	dialect := goqu.Dialect("postgres")
	var conninfo string = "host=localhost port=55000 user=postgres password=postgrespw dbname=wiz sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}
	db1 := dialect.DB(db)

	val, err:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).Insert(code,name)
	
	return val, err
}

func UpdateValue(code , name string)(string , error){
	dialect := goqu.Dialect("postgres")
	var conninfo string = "host=localhost port=55000 user=postgres password=postgrespw dbname=wiz sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}
	db1 := dialect.DB(db)
	val, err:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).Update(code,&name)
	
	return val, err

}

func GetOneValue(code string)(repositories.Currencies,bool,error,){

dialect := goqu.Dialect("postgres")
	var conninfo string = "host=localhost port=55000 user=postgres password=postgrespw dbname=wiz sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}
	db1 := dialect.DB(db)

	val, val2,err:=pgdb.NewCuurenciesRepository(*goqu.New(db1.Dialect(),db1.Db)).ReadOne(code)
	
	return val, val2,err
	}

func GetAllValue(pageNumber *uint,
	itemsPerPage uint,)(){}

func RemoveValue()(){}