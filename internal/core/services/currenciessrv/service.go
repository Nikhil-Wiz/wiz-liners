package currenciessrv

import (
	"fmt"
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/core/domain/services"
	"wiz-liners/internal/core/ports"
)
type service struct {
	currenciesRepository ports.CurrenciesRepository
}

func New(
	currenciesRepository ports.CurrenciesRepository,	
) *service {
	return &service{
		currenciesRepository: currenciesRepository,
	}
}

func (s *service) Create(
	Code string,
	Name string,
) (services.Currencies, error) {
	currencyCode, err := s.currenciesRepository.Insert(
		Code,
		Name,
	)
	if err != nil {
		return services.Currencies{}, err
	}

	repoCurrency, exists, err := s.currenciesRepository.ReadOne(currencyCode)

	if err != nil {
		return services.Currencies{}, err
	}

	if !exists {
		return services.Currencies{}, fmt.Errorf(
			"could not find Currency  with code: %s",
			currencyCode,
		)
	}

	return s.currencyRepoToService(
		repoCurrency,
	), nil
}

func (s *service) Get(code string) (services.Currencies, error) {
	repoCurrency, exists, err := s.currenciesRepository.ReadOne(code)

	if err != nil {
		return services.Currencies{}, err
	}

	if !exists {
		return services.Currencies{}, fmt.Errorf(
			"no currency found for code: %s",
			code,
		)
	}
	return s.currencyRepoToService(
		repoCurrency,
	), nil
}


func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]services.Currencies, error) {
	repoCurrencies, err := s.currenciesRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)

	if err != nil {
		return make([]services.Currencies, 0), err
	}

	currencies := []services.Currencies{}

	for _, cr := range repoCurrencies {
		currencies = append(
			currencies,
			s.currencyRepoToService(
				cr,
			),
		)
	}

	return currencies, nil

}

func (s *service) Modify(
	code string,
	Name *string,
) (services.Currencies, error) {
	_, err := s.currenciesRepository.Update(
		code,
		Name,
	)

	if err != nil {
		return services.Currencies{}, err
	}

	return s.Get(code)
}

func (s *service) Remove(code string) error {
	return s.currenciesRepository.Delete(code)
}


func (s *service) currencyRepoToService(
	repoCurrency repositories.Currencies,
) services.Currencies {
	return services.Currencies{
		Code:        repoCurrency.Code,
		Name:        repoCurrency.Name,
	}
}

