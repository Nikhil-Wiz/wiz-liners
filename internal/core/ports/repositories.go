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

	)()
	ReadOne(

	)()
	ReadMany(

	)()
	Update(

	)()
	Delete(

	)()

}

type CitiesRepository interface{
	Insert(

	)()
	ReadOne(

	)()
	ReadMany(

	)()
	Update(

	)()
	Delete(

	)()
	
}

type CurrenciesRepository interface{
	Insert(

	)()
	ReadOne(

	)()
	ReadMany(

	)()
	Update(

	)()
	Delete(

	)()
	
}

type Hs_CodeRepository interface{
	Insert(

	)()
	ReadOne(

	)()
	ReadMany(

	)()
	Update(

	)()
	Delete(

	)()
	
}

type PortsRepository interface{
	Insert(

	)()
	ReadOne(

	)()
	ReadMany(

	)()
	Update(

	)()
	Delete(

	)()
	
}