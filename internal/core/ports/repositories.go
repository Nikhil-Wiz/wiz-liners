package ports

import(
	"wiz-liners/internal/core/domain/repositories"
)

type LinersRepository interface{
	Insert(
		Name string,
		Code string,
		Type string,
		Logo string,
	)(int64,error)

	ReadOne(
		id int64,
	)(repositories.Liners,bool,error)
	
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Liners,error)

	Update(
		id int64,
		Name *string,
		Code *string,
		Type *string,
		Logo *string,
	)(int64,error)

	Delete(
		id int64,
	)(error)

}

type CountriesRepository interface{
	Insert(
		Name string,
		Iso_code string,
		Currency_code string,
	)(int64, error)

	ReadOne(
		id int64,
	)(repositories.Countries, bool, error)

	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Countries,error)

	Update(
		id int64,
		Name *string,
		Iso_code *string,
		Currency_code *string,
	)(int64, error)

	Delete(
		id int64,
	)(error)

}

type CitiesRepository interface{
	Insert(
		Name string,
		Country_id int64,
	)(int64, error)

	ReadOne(
		id int64,
	)(repositories.Cities, bool, error)

	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Cities,error)

	Update(
		id int64,
		Name *string,
		Country_id *int64,
	)(int64, error)

	Delete(
		id int64,
	)(error)
	
}

type CurrenciesRepository interface{
	Insert(
		Code string,
		Name string,
	)(repositories.Currencies, error)

	ReadOne(
		code string,
	)(repositories.Currencies, bool, error)

	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Currencies, error)

	Update(
		code *string,
		Name *string,
	)(repositories.Currencies, error)

	Delete(
		code string,
	)(error)
	
}

type Hs_CodeRepository interface{
	Insert(
		Code string, 
		Name string,
		Description string,
		Parent_code string,
	)(repositories.Hs_codes, error)

	ReadOne(
		code string,
	)(repositories.Hs_codes,bool, error)

	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Hs_codes, error)

	Update(
		code *string,
		Name *string,
		Description *string,
		Parent_code *string,
	)(repositories.Hs_codes, error)

	Delete(
		code string,
	)(error)
	
}

type PortsRepository interface{
	Insert(
		Code string, 
		Name string,
		Type string,
		City_Id int64,
		State string,
		Latitude float64,
		Longitude float64,
	)(repositories.Ports, error)

	ReadOne(
		code string,
	)(repositories.Ports, bool, error)

	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	)([]repositories.Ports, error)

	Update(
		code string,
		Name *string,
		Type *string,
		City_Id *int64,
		State *string,
		Latitude *float64,
		Longitude *float64,
	)(repositories.Ports, error)

	Delete(
		code string,
	)(error)
	
}