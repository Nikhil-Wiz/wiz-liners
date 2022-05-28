package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb/tables"
	hs_code "wiz-liners/internal/repositories/pgdb/tables/hs_codes"

	"github.com/doug-martin/goqu/v9"
)

type hscodeRepository struct {
	goquDB goqu.Database
}

func newHscodeRepository(goquDB goqu.Database) *hscodeRepository {
	return &hscodeRepository{
		goquDB: goquDB,
	}
}

func (r *hscodeRepository) Insert(
	code string,
	name string,
	description string,
	parent_code string,
) (string, error) {

	result, err := r.goquDB.Insert(tables.HS_CODES).Prepared(true).Rows(
		goqu.Record{
			hs_code.CODE: code,
			hs_code.NAME: name,
			hs_code.DESCRIPTION:description,
			hs_code.PARENT_CODE:parent_code,
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

func (r *hscodeRepository) ReadOne(
	code string,
) (repositories.Hs_codes, bool, error) {

	var c repositories.Hs_codes

	found, err := r.goquDB.From(
		tables.HS_CODES,
	).Prepared(true).Select(
		hs_code.CODE,
		hs_code.NAME,
		hs_code.DESCRIPTION,
		hs_code.PARENT_CODE,
	).Where(
		goqu.C(hs_code.CODE).Eq(code),
	).ScanStruct(&c)

	if err != nil {
		return repositories.Hs_codes{}, false, err
	}

	return c, found, nil
}

func (r *hscodeRepository) ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]repositories.Hs_codes, error) {

	var pgNum uint = 1

	if pageNumber == nil {
		pgNum = 1
	} else {
		pgNum = *pageNumber
	}

	if pgNum < 1 {
		pgNum = 1
	}

	var hscodeList []repositories.Hs_codes

	err := r.goquDB.From(
		tables.HS_CODES,
	).Prepared(true).Select(
		hs_code.CODE,
		hs_code.NAME,
		hs_code.DESCRIPTION,
		hs_code.PARENT_CODE,
	).Limit(itemsPerPage).Offset(
		(pgNum - 1) * itemsPerPage,
	).Order(
		goqu.I(hs_code.NAME).Asc(),
	).ScanStructs(&hscodeList)

	if err != nil {
		return []repositories.Hs_codes{}, err
	}

	return hscodeList, nil
}

func (r *hscodeRepository) Update(
	code string,
	Name *string,
	Description *string,
	Parent_code *string,
) (string, error) {

	record := goqu.Record{}

	if len(code) != 0 {
		record[hs_code.CODE] = *&code
	}

	if Name != nil {
		record[hs_code.NAME] = *Name
	}

	if Description != nil {
		record[hs_code.DESCRIPTION] = *Description
	}

	if Parent_code != nil {
		record[hs_code.PARENT_CODE] = *Parent_code
	}

	if len(record) == 0 {
		return "0", nil
	}

	result, err := r.goquDB.From(
		tables.HS_CODES,
	).Prepared(true).Update().Set(record).Where(
		goqu.C(hs_code.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return "0", nil
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return "0", nil
	}

	return string(affectedRows), nil
}

func (r *hscodeRepository) Delete(
	code string,
) error {

	_, err := r.goquDB.From(
		tables.HS_CODES,
	).Prepared(true).Delete().Where(
		goqu.C(hs_code.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}