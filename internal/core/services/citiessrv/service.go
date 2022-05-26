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
	currenciesRepository ports.CurrenciesRepository
	countriesRepository ports.CountriesRepository
}

func New(
	citiesRepository ports.CitiesRepository,
	currenciesRepository ports.CurrenciesRepository,
	countriesRepository ports.CountriesRepository,
	) *service {
	return &service{
		citiesRepository: citiesRepository,
		currenciesRepository: currenciesRepository,
		countriesRepository: countriesRepository,
	}
}

func (s *service) Create(
	id int64,
	name string,
	country_id int64,
) (services.Cities, error) {

	//repo for countries

	repoCountries , exists , err := s.countriesRepository.ReadOne(country_id)
	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"Country with id does not exists!:%d",
			id,
		)
	}

	//repo for currencies
	repoCurrencies , exists , err := s.currenciesRepository.ReadOne(repoCountries.Currency_code)
	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"Currency with id does not exists!:%d",
			id,
		)
	}

	repoCities, exists, err := s.citiesRepository.ReadOne(id)

	if err != nil {
		return services.Cities{}, err
	}
	if exists {
		return services.Cities{}, fmt.Errorf(
			"City with id exists!:%d",
			id,
		)
	}
	Id, err := s.citiesRepository.Insert(
		name,
		country_id,
	)
	log.Println(Id)
	if err != nil {
		return services.Cities{}, err
	}
	return s.citiesRepoToService(repoCities,repoCountries,repoCurrencies), nil
}

func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) (
	[]services.Cities, error,
) {
	repoCities, err := s.citiesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)

	if err != nil {
		return make([]services.Cities, 0), err
	}

	citiesIdsSet := mapset.NewSet[int64]()

	lo.ForEach(repoCities, func(rc repositories.Cities, _ int) {
		citiesIdsSet.Add(rc.Country_id)
	})

	citiesIds := citiesIdsSet.ToSlice()

	RepoCities, err := s.citiesRepository.ReadManyByIds(citiesIds)

	log.Println(RepoCities)

	if err != nil {
		return make([]services.Cities, 0), err
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

	return cities, err

}

func (s *service) Get(id int64) (services.Cities, error) {

	repoCities, exists, err := s.citiesRepository.ReadOne(id)

	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"No City with id exists!:%d",
			id,
		)
	}

	return s.citiesRepoToService(repoCities), nil

}
func (s *service) Modify(
	id int64,
	Name *string,
	Country_id *int64,
) (
	services.Cities, error) {

}
func (s *service) Remove(id int64) error {
	return s.citiesRepository.Delete(id)
}

func (s *service) citiesRepoToService(
	repoCities repositories.Cities,
	repoCountry repositories.Countries,
	repoCurrencies repositories.Currencies,
) services.Cities {
	return services.Cities{
		Id:   repoCities.Id,
		Name: repoCities.Name,
		Country_id: services.Countries{
			Id:       repoCountry.Id,
			Name:     repoCountry.Name,
			Iso_code: repoCountry.Iso_code,
			Currency_code: services.Currencies{
				Code: repoCurrencies.Code,
				Name: repoCurrencies.Name,
			},
		},
	}
}
