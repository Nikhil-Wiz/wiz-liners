package countriessrv

import (
	"fmt"
	"wiz-liners/internal/core/domain/services"
	"wiz-liners/internal/core/ports"
)

type service struct {
	countryRepository ports.CountriesRepository
	currencyRepository ports.CurrenciesRepository
}

func New(
	countryRepository ports.CountriesRepository,
	currencyRepository ports.CurrenciesRepository,
) *service {
	return &service{
		countryRepository: countryRepository,
		currencyRepository: currencyRepository,
	}
}

func (s *service) Create(
	id int64,
	name string,
	iso_code string,
	currency_code string,
) (services.Countries, error) {

	repoCountry, exists, err := s.countryRepository.ReadOne(id)

	if err != nil {
		return services.Countries{}, err
	}

	if !exists {
		return services.Countries{}, fmt.Errorf(
			"could not find the country with id: %d",
			id,
		)
	}

	
}