package citiessrv

import (
	"fmt"
	"log"
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/core/domain/services"
	"wiz-liners/internal/core/ports"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

type service struct {
	citiesRepository ports.CitiesRepository
}

func New(citiesRepository ports.CitiesRepository)*service{
	return &service{citiesRepository: citiesRepository}
}

func (s * service)Create(
	id int64,
	name string ,
	country_id int64,
)(services.Cities,error){
	repoCities,exists,err:=s.citiesRepository.ReadOne(id)

	if err != nil{
		return services.Cities{},err
	}
	if !exists{
		return services.Cities{},fmt.Errorf(
			"City with id exists!:%d",
			id,
		)
	}
	Id,err := s.citiesRepository.Insert(
		name,
		country_id,
	)
	log.Println(Id)
	if err != nil {
		return services.Cities{},err
	}
	return s.citiesRepoToService(repoCities),nil
}

func (s * service)GetMany(
	pageNumber *uint,
	itemsPerPage uint,
)(
	[]services.Cities,error,
){
	repoCities , err := s.citiesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)

	if err != nil {
		return make([]services.Cities,0),err
	}

	citiesIdsSet := mapset.NewSet[int64]()

	lo.ForEach(repoCities,func(rc repositories.Cities,_ int) {
		citiesIdsSet.Add(rc.Country_id)
	})

	citiesIds := citiesIdsSet.ToSlice()

	RepoCities , err :=  s.citiesRepository.ReadManyByIds(citiesIds)

	log.Println(RepoCities)

	if err != nil {
		return make([]services.Cities,0),err
	}

	citiesIdMap := map[int64]repositories.Cities{}

	cities := []services.Cities{}

	for _, cr := range repoCities {
		_, exists := citiesIdMap[cr.Id]

		if !exists {

			return make([]services.Cities, 0), fmt.Errorf(
				"currency not found against id: %d",
				cr.Id,
			)
		}

		cities = append(
			cities,
			s.citiesRepoToService(
				cr,
			),
		)
	}

	return cities,err

}

func (s * service)Get(id int64)(services.Cities,error){

	repoCities , exists , err := s.citiesRepository.ReadOne(id)

	if err != nil{
		return services.Cities{},err
	}
	if !exists{
		return services.Cities{},fmt.Errorf(
			"No City with id exists!:%d",
			id,
		)
	}

	return s.citiesRepoToService(repoCities),nil


}
func (s * service)Modify(
	id int64,
	Name *string,
	Country_id *int64,
)(
	services.Cities,error){
		


}
func (s * service)Remove(id int64)(error){
	return s.citiesRepository.Delete(id)
}

func (s *service) citiesRepoToService(
	repoCities repositories.Cities,
) services.Cities {
	return services.Cities{
		Id:      repoCities.Id,
		Name:    repoCities.Name,
		Country_id: repoCities.Country_id,
	}
}