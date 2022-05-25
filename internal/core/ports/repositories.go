package ports

import(
	"wiz-liners/internal/core/domain/repositories"
)

type LinerRepository interface{
	Insert(
	Name string,
	Code string,
	Logo string,
	TypeOfLiners string,
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
		Logo *string,
		TypeOfLiners *string,
	)(int64,error)

	Delete(
		id int64,
	)(error)

}