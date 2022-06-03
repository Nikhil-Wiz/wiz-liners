package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb/tables"
	"wiz-liners/internal/repositories/pgdb/tables/currencies"

	"github.com/doug-martin/goqu/v9"
)

type currenciesRepository struct {
	goquDB goqu.Database
}

func NewCuurenciesRepository(goquDB goqu.Database) *currenciesRepository {
	return &currenciesRepository{
		goquDB: goquDB,
	}
}

func (r *currenciesRepository) Insert(
	code string,
	name string,
) (
	string, error) {
	result, err := r.goquDB.Insert(tables.CURRENCIES).Rows(
		goqu.Record{
			currencies.CODE:        code,
			currencies.NAME:        name,
		},
	).Executor().Exec()

	if err != nil {
		return "0", err
	}

	rowId, err := result.LastInsertId()

	if err != nil {
		return "0", err
	}

	return string(rune(rowId)), nil
}

func (r *currenciesRepository) ReadOne(
	code string,
) (
	repositories.Currencies,
	bool,
	error,
) {
	var c repositories.Currencies

	found, err := r.goquDB.From(
		tables.CURRENCIES,
	).Select(
		currencies.CODE,
		currencies.NAME,
	).Where(
		goqu.C(currencies.CODE).Eq(code),
	).ScanStruct(&c)

	if err != nil {
		return repositories.Currencies{}, false, err
	}

	return c, found, nil
}
func (r *currenciesRepository) ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
) (
	[]repositories.Currencies,
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

	var currenciesList []repositories.Currencies

	err := r.goquDB.From(
		tables.CURRENCIES,
	).Select(
		currencies.CODE,
		currencies.NAME,
		currencies.CREATED_AT,
		currencies.MODIFIED_AT,
	).Limit(itemsPerPage).Offset(
		(pn - 1) * itemsPerPage,
	).Order(
		goqu.I(currencies.NAME).Asc(),
	).ScanStructs(&currenciesList)

	if err != nil {
		return []repositories.Currencies{}, err
	}

	return currenciesList, nil
}
func (r *currenciesRepository) Update(
	code string,
	name *string,

) (
	string,
	error,
) {
	record := goqu.Record{}

	if len(code) != 0 {
		record[currencies.CODE] = code
	}

	if name != nil {
		record[currencies.NAME] = *name
	}

	if len(record) == 0 {
		return "0", nil
	}

	result, err := r.goquDB.From(
		tables.CURRENCIES,
	).Update().Set(record).Where(
		goqu.C(currencies.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return "0", err
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return "0", err
	}

	return string(rune(affectedRows)), nil
}
func (r *currenciesRepository) Delete(
	code string,
) error {
	_, err := r.goquDB.From(
		tables.CURRENCIES,
	).Delete().Where(
		goqu.C(currencies.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}
