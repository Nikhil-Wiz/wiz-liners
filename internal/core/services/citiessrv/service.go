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
	citiesRepository     ports.CitiesRepository
	currenciesRepository ports.CurrenciesRepository
	countriesRepository  ports.CountriesRepository
}

func New(
	citiesRepository ports.CitiesRepository,
	currenciesRepository ports.CurrenciesRepository,
	countriesRepository ports.CountriesRepository,
) *service {
	return &service{
		citiesRepository:     citiesRepository,
		currenciesRepository: currenciesRepository,
		countriesRepository:  countriesRepository,
	}
}

func (s *service) Create(
	id int64,
	name string,
	country_id int64,
) (services.Cities, error) {

	//repo for countries

	repoCountries, exists, err := s.countriesRepository.ReadOne(country_id)
	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"country with id does not exists!:%d",
			id,
		)
	}

	//repo for currencies
	repoCurrencies, exists, err := s.currenciesRepository.ReadOne(repoCountries.Currency_code)
	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"currency with id does not exists!:%d",
			id,
		)
	}

	repoCities, exists, err := s.citiesRepository.ReadOne(id)

	if err != nil {
		return services.Cities{}, err
	}
	if exists {
		return services.Cities{}, fmt.Errorf(
			"city with id exists!:%d",
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
	return s.citiesRepoToService(repoCities, repoCountries, repoCurrencies), nil
}

func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) (
	[]services.Cities, error,
) {

	repoCountries, err := s.countriesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)
	if err != nil {
		return []services.Cities{}, err
	}

	//repo for currencies
	repoCurrencies, err := s.currenciesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)
	
	if err != nil {
		return make([]services.Cities, 0), err
	}

	currencyIdsSet := mapset.NewSet[int64]()

	lo.ForEach(repoCountries, func(rc repositories.Countries, _ int) {
		currencyIdsSet.Add(rc.Id)
	})

	countryIdsSet := mapset.NewSet[int64]()

	//currencyIds := currencyIdsSet.ToSlice()
	countriesIds := countryIdsSet.ToSlice()

	repoCities, err := s.citiesRepository.ReadManyByIds(
		countriesIds,
	)
	if err != nil {
		return []services.Cities{}, err
	}
	lo.ForEach(repoCities, func(rc repositories.Cities, _ int) {
		countryIdsSet.Add(rc.Country_id)
	})

	currenciesIdMap := map[string]repositories.Currencies{}
	countriesIdMap := map[int64]repositories.Countries{}

	lo.ForEach(repoCurrencies, func(cr repositories.Currencies, _ int) {
		currenciesIdMap[cr.Code] = cr
	})
	lo.ForEach(repoCountries, func(cr repositories.Countries, _ int) {
		countriesIdMap[cr.Id] = cr
	})

	cities := []services.Cities{}

	for _, cr := range repoCities {
		countriessRepository := countriesIdMap[cr.Country_id]
		currenciesRepository, exists := currenciesIdMap[countriessRepository.Currency_code]

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
				countriessRepository,
				currenciesRepository,
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
			"no City with id exists!:%d",
			id,
		)
	}
	repoCountries, exists, err := s.countriesRepository.ReadOne(repoCities.Country_id)

	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"no Country with id exists!:%d",
			id,
		)
	}
	repoCurrencies, exists, err := s.currenciesRepository.ReadOne(repoCountries.Currency_code)

	if err != nil {
		return services.Cities{}, err
	}
	if !exists {
		return services.Cities{}, fmt.Errorf(
			"no Currency with id exists!:%d",
			id,
		)
	}

	return s.citiesRepoToService(repoCities, repoCountries, repoCurrencies), nil

}
func (s *service) Modify(
	id int64,
	Name *string,
	Country_id *int64,
) (
	services.Cities, error) {

	if Country_id != nil {
		_, exists, err := s.citiesRepository.ReadOne(*Country_id)

		if !exists {
			return services.Cities{}, fmt.Errorf(
				"no Currency with id exists!:%d",
				id,
			)
		}
		if err != nil {
			return services.Cities{}, err
		}
	}
	_, err := s.citiesRepository.Update(
		id,
		Name,
		Country_id,
	)
	if err != nil {
		return services.Cities{}, err
	}

	return s.Get(id)

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
