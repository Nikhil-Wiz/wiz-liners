package hscodessrv

import (
	"fmt"
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/core/domain/services"
	"wiz-liners/internal/core/ports"
)

type service struct {
	Hs_CodeRepository ports.Hs_CodeRepository
}

func New(
	Hs_CodeRepository ports.Hs_CodeRepository,
) *service {
	return &service{
		Hs_CodeRepository: Hs_CodeRepository,
	}
}

func (s *service) Create(
	Code string,
	Name string,
	Description string,
	Parent_code string,
) (services.Hs_codes, error) {
	hsCode, err := s.Hs_CodeRepository.Insert(
		Code,
		Name,
		Description,
		Parent_code,
	)
	if err != nil {
		return services.Hs_codes{}, err
	}

	repoHsCode, exists, err := s.Hs_CodeRepository.ReadOne(hsCode)

	if err != nil {
		return services.Hs_codes{}, err
	}

	if !exists {
		return services.Hs_codes{}, fmt.Errorf(
			"could not find HS Code with code: %s",
			hsCode,
		)
	}

	return s.hsCodeRepoToService(
		repoHsCode,
	), nil
}

func (s *service) Get(code string) (services.Hs_codes, error) {
	repoHsCode, exists, err := s.Hs_CodeRepository.ReadOne(code)

	if err != nil {
		return services.Hs_codes{}, err
	}

	if !exists {
		return services.Hs_codes{}, fmt.Errorf(
			"no HS Code found for code: %s",
			code,
		)
	}
	return s.hsCodeRepoToService(
		repoHsCode,
	), nil
}

func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]services.Hs_codes, error) {
	repoHsCodes, err := s.Hs_CodeRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)

	if err != nil {
		return make([]services.Hs_codes, 0), err
	}

	hsCodes := []services.Hs_codes{}

	for _, cr := range repoHsCodes {
		hsCodes = append(
			hsCodes,
			s.hsCodeRepoToService(
				cr,
			),
		)
	}

	return hsCodes, nil

}

func (s *service) Modify(
	code string,
	Name *string,
	Description *string,
	Parent_code *string,
) (services.Hs_codes, error) {
	_, err := s.Hs_CodeRepository.Update(
		code,
		Name,
		Description,
		Parent_code,
	)

	if err != nil {
		return services.Hs_codes{}, err
	}

	return s.Get(code)
}

func (s *service) Remove(code string) error {
	return s.Hs_CodeRepository.Delete(code)
}

func (s *service) hsCodeRepoToService(
	repoHsCode repositories.Hs_codes,
) services.Hs_codes {
	return services.Hs_codes{
		Code:        repoHsCode.Code,
		Name:        repoHsCode.Name,
		Description: repoHsCode.Description,
		Parent_code: repoHsCode.Parent_code,
	}
}
