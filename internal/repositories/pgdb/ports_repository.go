package pgdb

import (
	"wiz-liners/internal/core/domain/repositories"
	//"wiz-liners/internal/core/ports"
	"wiz-liners/internal/repositories/pgdb/tables"
	"wiz-liners/internal/repositories/pgdb/tables/ports"

	"github.com/doug-martin/goqu/v9"
)

type portsRepository struct {
	goquDB goqu.Database
}

func newPortsRepository(goquDB goqu.Database) *portsRepository {
	return &portsRepository{
		goquDB: goquDB,
	}
}

func (r *portsRepository) Insert(
	code string,
	name string,
	Type string,
	city_Id int64,
	state string,
	latitude float64,
	longitude float64,
) (
	string,
	error) {
	result, err := r.goquDB.Insert(tables.PORTS).Prepared(true).Rows(
		goqu.Record{
			ports.CODE:      code,
			ports.NAME:      name,
			ports.TYPE:      Type,
			ports.CITYID:    city_Id,
			ports.STATE:     state,
			ports.LATITUDE:  latitude,
			ports.LONGITUDE: longitude,
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

func (r *portsRepository) ReadOne(
	code string,
) (
	repositories.Ports,
	bool,
	error,
) {
	var c repositories.Ports

	found, err := r.goquDB.From(
		tables.PORTS,
	).Prepared(true).Select(
			ports.CODE,
			ports.NAME,
			ports.TYPE,
			ports.CITYID,
			ports.STATE,
			ports.LATITUDE,
			ports.LONGITUDE,
			ports.CREATED_AT,
			ports.MODIFIED_AT,
	).Where(
		goqu.C(ports.CODE).Eq(code),
	).ScanStruct(&c)

	if err != nil {
		return repositories.Ports{}, false, err
	}

	return c, found, nil
}
func (r *portsRepository) ReadMany(
	pageNumber *uint,
	itemsPerPage uint,
) (
	[]repositories.Ports,
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

	var portsList []repositories.Ports

	err := r.goquDB.From(
		tables.CURRENCIES,
	).Prepared(true).Select(
			ports.CODE,
			ports.NAME,
			ports.TYPE,
			ports.CITYID,
			ports.STATE,
			ports.LATITUDE,
			ports.LONGITUDE,
			ports.CREATED_AT,
			ports.MODIFIED_AT,
	).Limit(itemsPerPage).Offset(
		(pn - 1) * itemsPerPage,
	).Order(
		goqu.I(ports.NAME).Asc(),
	).ScanStructs(&portsList)

	if err != nil {
		return []repositories.Ports{}, err
	}

	return portsList, nil
}
func (r *portsRepository) Update(
		code string,
		Name *string,
		Type *string,
		City_Id *int64,
		State *string,
		Latitude *float64,
		Longitude *float64,
) (
	string,
	error,
) {
	record := goqu.Record{}

	if len(code) != 0 {
		record[ports.CODE] = *&code
	}

	if Name != nil {
		record[ports.NAME] = *Name
	}

	if Type != nil {
		record[ports.TYPE] = *Type
	}
	if City_Id != nil {
		record[ports.CITYID] = *City_Id
	}
	if State != nil {
		record[ports.STATE] = *State
	}
	if Latitude != nil {
		record[ports.LATITUDE] = *Latitude
	}
	if Longitude != nil {
		record[ports.LONGITUDE] = *Longitude
	}

	if len(record) == 0 {
		return "0", nil
	}

	result, err := r.goquDB.From(
		tables.PORTS,
	).Prepared(true).Update().Set(record).Where(
		goqu.C(ports.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return "0", err
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return "0", err
	}

	return string(affectedRows), nil
}
func (r *portsRepository) Delete(
	code string,
) error {
	_, err := r.goquDB.From(
		tables.PORTS,
	).Prepared(true).Delete().Where(
		goqu.C(ports.CODE).Eq(code),
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}
