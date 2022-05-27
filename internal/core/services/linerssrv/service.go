
package linerssrv

import (
	"fmt"
	"log"
	"wiz-liners/internal/core/domain/repositories"
	"wiz-liners/internal/core/domain/services"
	"wiz-liners/internal/core/ports"
)

type service struct {
	linersRepository ports.LinersRepository
}

func New(
	linersRepository ports.LinersRepository,
) *service {
	return &service{
		linersRepository: linersRepository,
	}
}

func (s *service) Create(
	id int64,
	name string,
	code string,
	Type string,
	logo string,
) (services.Liners, error) {

	repoLiner, exists, err := s.linersRepository.ReadOne(id)

	if err != nil {
		return services.Liners{}, err
	}

	if !exists {
		return services.Liners{}, fmt.Errorf(
			"could not find the liner with id: %d",
			id,
		)
	}

	Id, err := s.linersRepository.Insert(
		name,
		code,
		Type,
		logo,
	)
	log.Println(Id)

	if err != nil {
		return services.Liners{}, err
	}

	return s.linersRepoToService(
		repoLiner,
	), nil
}

func (s *service) Get(id int64) (services.Liners, error) {

	repoLiners, exists, err := s.linersRepository.ReadOne(id)

	if err != nil {
		return services.Liners{}, err
	}

	if !exists {
		return services.Liners{}, fmt.Errorf(
			"no liner found for id: %d",
			id,
		)
	}

	return s.linersRepoToService(
		repoLiners,
	), nil
}

func (s *service) Modify(
	id int64,
	Name *string,
	Code *string,
	Type *string,
	Logo *string,
) (services.Liners, error) {

	if id != 0 {
		_, exists, err := s.linersRepository.ReadOne(id)

		if err != nil {
			return services.Liners{}, err
		}

		if !exists {
			return services.Liners{}, fmt.Errorf(
				"could not find liner with id: %d",
				id,
			)
		}
	}

	_, err := s.linersRepository.Update(
		id,
		Name,
		Code,
		Type,
		Logo,
	)

	if err != nil {
		return services.Liners{}, err
	}

	return s.Get(id)
}

func (s *service) GetMany(
	pageNumber *uint,
	itemsPerPage uint,
) ([]services.Liners, error) {

	repoLiners, err := s.linersRepository.ReadMany(
		pageNumber,
		itemsPerPage,
	)

	if err != nil {
		return make([]services.Liners, 0), err
	}

	liners := []services.Liners{}

	for _, cr := range repoLiners {
		liners = append(
			liners, 
			s.linersRepoToService(
				cr,
			),
		)
	}

	return liners, nil
}

func (s *service) Remove(id int64) error {
	return s.linersRepository.Delete(id)
}

func (s *service) linersRepoToService(
	repoLiner repositories.Liners,
) services.Liners {
	return services.Liners{
		Id: repoLiner.Id,
		Name: repoLiner.Name,
		Code: repoLiner.Code,
		Type: repoLiner.Type,
		Logo: repoLiner.Logo,
	}
}
