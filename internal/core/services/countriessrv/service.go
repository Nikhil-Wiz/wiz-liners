package countriessrv

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
	currenciesRepository ports.CurrenciesRepository
	countriesRepository ports.CountriesRepository
}

func New(
	currenciesRepository ports.CurrenciesRepository,
	countriesRepository ports.CountriesRepository,
) *service {
	return &service{
		currenciesRepository: currenciesRepository,
		countriesRepository: countriesRepository,
	}
}

func (s *service) Create(
	id int64,
	name string,
	iso_code string,
	currency_code string,
) (services.Countries, error) {

	//repo for countries
	repoCountries, exists, err := s.countriesRepository.ReadOne(id)

	if err != nil {
		return services.Countries{}, err
	}

	if !exists {
		return services.Countries{}, fmt.Errorf(
			"not found country with id: %d",
			id,
		)
	}

	//repo for currencies
	repoCurrencies, exists, err := s.currenciesRepository.ReadOne(currency_code)

	if err != nil {
		return services.Countries{}, err
	}

	if !exists {
		return services.Countries{}, fmt.Errorf(
			"currency code does not exists!:%s",
			currency_code,
		)
	}

	Id, err := s.countriesRepository.Insert( 
		name,
		iso_code,
		currency_code,
	)

	log.Println(Id)

	if err != nil {
		return services.Countries{}, err
	}

	return s.countriesRepoToService(repoCountries, repoCurrencies), nil
}

func (s *service) Get(id int64) (services.Countries, error){

	repoCountries, exists, err := s.countriesRepository.ReadOne(id)

	if err != nil {
		return services.Countries{}, err
	}

	if !exists {
		return services.Countries{}, fmt.Errorf(
			"no country with id:%d",
			id,
		)
	}

	repoCurrencies, exists, err := s.currenciesRepository.ReadOne(repoCountries.Currency_code)

	if err != nil {
		return services.Countries{}, err
	}

	if !exists {
		return services.Countries{}, fmt.Errorf(
			"no currency with code:%s",
			repoCountries.Currency_code,
		)
	}

	return s.countriesRepoToService(repoCountries, repoCurrencies), nil
}

func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]services.Countries, error) {

	repoCountries, err := s.countriesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)
	if err != nil {
		return []services.Countries{}, err
	}

	//repo for currencies
	currencyIdsSet := mapset.NewSet[string]()
	lo.ForEach(repoCountries, func(rc repositories.Countries, _ int) {
		currencyIdsSet.Add(rc.Currency_code)
	})
	currencyIds := currencyIdsSet.ToSlice()

	repoCurrencies, err := s.currenciesRepository.ReadManyByCode(
		currencyIds,
	)
	
	if err != nil {
		return make([]services.Countries, 0), err
	}

	currenciesIdMap := map[string]repositories.Currencies{}

	lo.ForEach(repoCurrencies, func(cr repositories.Currencies, _ int) {
		currenciesIdMap[cr.Code] = cr
	})

	countries := []services.Countries{}

	for _, cr := range repoCountries {
		currenciesRepository, exists := currenciesIdMap[cr.Currency_code]

		if !exists {

			return make([]services.Countries, 0), fmt.Errorf(
				"currency not found against id: %d",
				cr.Id,
			)
		}

		countries = append(
			countries,
			s.countriesRepoToService(
				cr,
				currenciesRepository,
			),
		)
	}

	return countries, err
}

func (s *service) Modify(
	id int64,
	Name *string,
	Iso_code *string,
	Currency_code *string,
) (services.Countries, error) {

	if Currency_code != nil {
		_, exists, err := s.countriesRepository.ReadOne(id)

		if err != nil {
			return services.Countries{}, err
		}

		if !exists {
			return services.Countries{}, fmt.Errorf(
				"no country with id exists!:%d",
				id,
			)
		}
	}

	_, err := s.countriesRepository.Update(
		id,
		Name,
		Iso_code,
		Currency_code,
	)

	if err != nil {
		return services.Countries{}, err
	}

	return s.Get(id)
}

func (s *service) Remove(id int64) error {
	return s.countriesRepository.Delete(id)
}

func (s *service) countriesRepoToService(
	repoCountries repositories.Countries,
	repoCurrencies repositories.Currencies,
) services.Countries {
	return services.Countries{
		Name: repoCountries.Name,
		Iso_code: repoCountries.Iso_code,
		Currency_code: services.Currencies{
			Code: repoCurrencies.Code,
			Name: repoCurrencies.Name,
		},
	}
}