package ports

import "wiz-liners/internal/core/domain/services"

type CitiesService interface{
	Create(

		)()
		GetMany(
	
		)()
		Get(
	
		)()
		Modify(
	
		)()
		Remove(
	
		)()
}
type CountriesService interface{
	Create(

		)()
		GetMany(
	
		)()
		Get(
	
		)()
		Modify(
	
		)()
		Remove(
	
		)()
	
}
type CurrenciesService interface{
	Create(

		)()
		GetMany(
	
		)()
		Get(
	
		)()
		Modify(
	
		)()
		Remove(
	
		)()
	
}
type Hs_CodeService interface{
	Create(

		)()
		GetMany(
	
		)()
		Get(
	
		)()
		Modify(
	
		)()
		Remove(
	
		)()
	
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
	
		remove(
			id int64,
		)(error)
	
}
type PortsService interface{
	Create(

		)()
		GetMany(
	
		)()
		Get(
	
		)()
		Modify(
	
		)()
		Remove(
	
		)()
	
}