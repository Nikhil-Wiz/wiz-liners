package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb/tables"
	"wiz-liners/internal/repositories/pgdb/tables/countries"

	"github.com/doug-martin/goqu/v9"
)

type countriesRepository struct {
	goquDB goqu.Database
}

func newCountriesRepository(goquDB goqu.Database) *countriesRepository {
	return &countriesRepository{
		goquDB: goquDB,
	}
}

func (r *countriesRepository) Insert(
		name string,
		iso_code string,
		currency_code string,
) (   int64,error) {
	result, err := r.goquDB.Insert(tables.COUNTRIES).Prepared(true).Rows(
		goqu.Record{
			countries.NAME:name,
			countries.ISO_CODE:iso_code,
			countries.CURRENCY_CODE:currency_code,
		},
	).Executor().Exec()

	if err != nil {
		return 0, err
	}

	rowId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return rowId, nil
}

func (r *countriesRepository) ReadOne(
	id int64,
) (
	repositories.Countries,
	bool,
	error,
) {
	var c repositories.Countries

	found, err := r.goquDB.From(
		tables.COUNTRIES,
	).Prepared(true).Select(
			countries.NAME,
			countries.ISO_CODE,
			countries.CURRENCY_CODE,
			countries.CREATED_AT,
			countries.MODIFIED_AT,
	).Where(
		goqu.C(countries.ID).Eq(id),
	).ScanStruct(&c)

	if err != nil {
		return repositories.Countries{}, false, err
	}

	return c, found, nil
}
func (r *countriesRepository) ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
) (
	[]repositories.Countries,
	error,
) {
	var pn uint = 1
	if pageNumber == nil {
		pn = 1
	} else {
		pn = *pageNumber
	}

	if pn < 1 {
		pn = 1
	}

	var countriesList []repositories.Countries

	err := r.goquDB.From(
		tables.COUNTRIES,
	).Prepared(true).Select(
		countries.NAME,
		countries.ISO_CODE,
		countries.CURRENCY_CODE,
		countries.CREATED_AT,
		countries.MODIFIED_AT,
	).Limit(itemsPerPage).Offset(
		(pn - 1) * itemsPerPage,
	).Order(
		goqu.I(countries.NAME).Asc(),
	).ScanStructs(&countriesList)

	if err != nil {
		return []repositories.Countries{}, err
	}

	return countriesList, nil
}
func (r *countriesRepository) Update(
	id int64,
	Name *string,
	Iso_code *string,
	Currency_code *string,
) (
	int64,
	error,
) {
	record := goqu.Record{}

	if id != 0 {
		record[countries.ID] = id
	}

	if Name != nil {
		record[countries.NAME] = *Name
	}

	if Iso_code != nil {
		record[countries.ISO_CODE] = *Iso_code
	}
	if Currency_code != nil {
		record[countries.ISO_CODE] = *Iso_code
	}
	

	if len(record) == 0 {
		return 0, nil
	}

	result, err := r.goquDB.From(
		tables.COUNTRIES,
	).Prepared(true).Update().Set(record).Where(
		goqu.C(countries.ID).Eq(id),
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
func (r *countriesRepository) Delete(
	id int64,
) error {
	_, err := r.goquDB.From(
		tables.COUNTRIES,
	).Prepared(true).Delete().Where(
		goqu.C(countries.ID).Eq(id),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}
