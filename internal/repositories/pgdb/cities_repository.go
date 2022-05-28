package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb/tables"
	"wiz-liners/internal/repositories/pgdb/tables/cities"

	"github.com/doug-martin/goqu/v9"
)

type citiesRepository struct {
	goquDB goqu.Database
}

func newCitiesRepository(goquDB goqu.Database) *citiesRepository {
	return &citiesRepository{
		goquDB: goquDB,
	}
}

func (r *citiesRepository) Insert(
	name string,
	country_id int64,
) (string, error) {

	result, err := r.goquDB.Insert(tables.CITIES).Prepared(true).Rows(
		goqu.Record{
			cities.NAME: name,
			cities.COUNTRY_ID: country_id,
		},
	).Executor().Exec()

	if err != nil {
		return "0", err
	}

	rowId, err := result.LastInsertId()

	if err != nil {
		return "0", err
	}

	return string(rowId), nil
}

func (r *citiesRepository) ReadOne(
	id int64,
) (repositories.Cities, bool, error) {

	var c repositories.Cities

	found, err := r.goquDB.From(
		tables.CITIES,
	).Prepared(true).Select(
		cities.ID,
		cities.NAME,
		cities.COUNTRY_ID,
	).Where(
		goqu.C(cities.ID).Eq(id),
	).ScanStruct(&c)


	if err != nil {
		return repositories.Cities{}, false, err
	}

	return c, found, err
}

func (r *citiesRepository) ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]repositories.Cities, error) {

	var pgNum uint = 1

	if pageNumber == nil {
		pgNum = 1
	} else {
		pgNum = *pageNumber
	}

	if pgNum < 1 {
		pgNum = 1
	}

	var citiesList []repositories.Cities

	err := r.goquDB.From(
		tables.CITIES,
	).Prepared(true).Select(
		cities.ID,
		cities.NAME,
		cities.COUNTRY_ID,
	).Limit(itemsPerPage).Offset(
		(pgNum - 1) * itemsPerPage,
	).Order(
		goqu.I(cities.NAME).Asc(),
	).ScanStructs(&citiesList)

	if err != nil {
		return []repositories.Cities{}, err
	}

	return citiesList, nil
}

func (r *citiesRepository) Update(
	id int64,
	name *string,
	country_id *int64,
) (int64, error) {

	record := goqu.Record{}

	if name != nil {
		record[cities.NAME] = *name
	}

	if country_id != nil {
		record[cities.COUNTRY_ID] = *country_id
	}

	if len(record) == 0 {
		return 0, nil
	}

	result, err := r.goquDB.From(
		tables.CITIES,
	).Prepared(true).Update().Set(record).Where(
		goqu.C(cities.ID).Eq(id),
	).Executor().Exec()

	if err != nil {
		return 0, err
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (r *citiesRepository) Delete(
	id int64,
) (error) {

	_, err := r.goquDB.From(
		tables.CITIES,
	).Prepared(true).Delete().Where(
		goqu.C(cities.ID).Eq(id),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}