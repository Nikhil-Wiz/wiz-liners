package ports

import "wiz-liners/internal/core/domain/services"

type CitiesService interface{
	Create(
		Name string,
		Country_Id int64,
	)(services.Cities,error)
	
	Get(
		id int64,
	)(services.Cities, error)

	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Cities, error)

	Modify(
		id int64,
		Name *string,
		Country_Id *int64,
	)(services.Cities, error)

	Remove(
		id int64,
	)(error)
}
type CountriesService interface{
	Create(
		Name string,
		Iso_code string,
		Currency_code string,
	)(services.Countries, error)

	Get(
		id int64,
	)(services.Countries, error)

	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Countries, error)

	Modify(
		id int64,
		Name *string,
		Iso_code *string,
		Currency_code *string,
	)(services.Countries, error)

	Remove(
		id int64,
	)(error)
	
}
type CurrenciesService interface{
	Create(
		Code string,
		Name string,
	)(services.Currencies, error)

	Get(
		code string,
	)(services.Currencies, error)

	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Currencies, error)

	Modify(
		code string,
		Name *string,
	)(services.Currencies, error)

	Remove(
		code string,
	)(error)
	
}
type Hs_CodeService interface{
	Create(
		Code string,
		Name string,
		Description string,
		Parent_code string,
	)(services.Hs_codes, error)

	Get(
		code string,
	)(services.Hs_codes, error)

	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Hs_codes, error)

	Modify(
		code string,
		Name *string,
		Description *string,
		Parent_code *string,
	)(services.Hs_codes, error)

	Remove(
		code string,
	)(error)

}
type LinersService interface{
	Create(
		Name string,
		Code string,
		Type string,
		Logo string,
	)(services.Liners,error)
	
	Get(
		id int64,
	)(services.Liners,error)
		
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Liners,error)
	
	Modify(
		id int64,
		Name *string,
		Code *string,
		Type *string,
		Logo *string,
	)(services.Liners,error)
	
	Remove(
		id int64,
	)(error)
	
}
type PortsService interface{
	Create(
		Code string,
		Name string,
		Type string,
		City_Id int64,
		State string,
		Latitude float64,
		Longitude float64,
	)(services.Port, error)
	
	Get(
		code string,
	)(services.Port, error)

	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]services.Port, error)

	Modify(
		code string,
		Name *string,
		Type *string,
		City_Id *int64,
		State *string,
		Latitude *float64,
		Longitude *float64,
	)(services.Port, error)

	Remove(
		code string,
	)(error)
	
}