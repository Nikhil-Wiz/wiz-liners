package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/repositories/pgdb/tables"
	"wiz-liners/internal/repositories/pgdb/tables/liners"

	goqu "github.com/doug-martin/goqu/v9"
)

type linersRepository struct {
	goquDB goqu.Database
}

func newLinersRepository(goquDB goqu.Database) *linersRepository {
	return &linersRepository{
		goquDB: goquDB,
	}
}

func (r *linersRepository)Insert(
	name string,
	code string,
	Type string, //sea or air
	logo string,
)(
	int64,error){
	result, err := r.goquDB.Insert(tables.LINERS).Prepared(true).Rows(
		goqu.Record{
			liners.NAME : name,
			liners.CODE : code,
			liners.TYPE : Type,
			liners.LOGO : logo,

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

func (r * linersRepository)ReadOne(
	id int64,
)(
	repositories.Liners,
	bool,
	error,
){
	var c repositories.Liners

	found, err := r.goquDB.From(
		tables.LINERS,
	).Prepared(true).Select(
		liners.ID,
		liners.NAME,
		liners.CODE,
		liners.TYPE,
		liners.LOGO,
		liners.CREATED_AT,
		liners.MODIFIED_AT,
	).Where(
		goqu.C(liners.ID).Eq(id),
	).ScanStruct(&c)

	if err != nil {
		return repositories.Liners{}, false, err
	}

	return c, found, nil
}
func (r * linersRepository)ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
)(
	[]repositories.Liners,
	error,
){
	var pn uint = 1
	if pageNumber == nil {
		pn = 1
	} else {
		pn = *pageNumber
	}

	if pn < 1 {
		pn = 1
	}

	var linersList []repositories.Liners

	err := r.goquDB.From(
		tables.LINERS,
	).Prepared(true).Select(
		liners.ID,
		liners.NAME,
		liners.CODE,
		liners.TYPE,
		liners.LOGO,
		liners.CREATED_AT,
		liners.MODIFIED_AT,
	).Limit(itemsPerPage).Offset(
		(pn - 1) * itemsPerPage,
	).Order(
		goqu.I(liners.NAME).Asc(),
	).ScanStructs(&linersList)

	if err != nil {
		return []repositories.Liners{}, err
	}

	return linersList, nil
}
func (r * linersRepository)Update(
	id int64,
		name *string,
		code *string,
		Type *string,
		logo *string,
)(
	int64,
	error,
){
	record := goqu.Record{}

	if name != nil {
		record[liners.NAME] = *name
	}

	if code != nil {
		record[liners.CODE] = *code
	}

	if Type != nil {
		record[liners.TYPE] = *Type
	}
	if logo != nil {
		record[liners.LOGO] = *logo
	}

	if len(record) == 0 {
		return 0, nil
	}

	result, err := r.goquDB.From(
		tables.LINERS,
	).Prepared(true).Update().Set(record).Where(
		goqu.C(liners.ID).Eq(id),
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
func (r * linersRepository)Delete(
	id int64,
)(
	error,
){
	_, err := r.goquDB.From(
		tables.LINERS,
	).Prepared(true).Delete().Where(
		goqu.C(liners.ID).Eq(id),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}