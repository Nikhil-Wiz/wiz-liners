package portssrv
import (
    "fmt"
    "log"
    "wiz-liners/internal/core/domain/repositories"
    "wiz-liners/internal/core/domain/services"
    "wiz-liners/internal/core/ports"
)
type service struct {
    portRepository ports.PortsRepository
}
func New(
    portRepository ports.PortsRepository,
) *service {
    return &service{
        portRepository: portRepository,
    }
}
func (s *service) Create(
    code string,
    name string,
    Type string,
    city_Id int64,
    state string,
    latitude float64,
    longitude float64,
) (services.Port, error) {
    repoPort, exists, err := s.portRepository.ReadOne(code)
    if err != nil {
        return services.Port{}, err
    }
    if !exists {
        return services.Port{}, fmt.Errorf(
            "could not find the port with code: %s",
            code,
        )
    }
    Code, err := s.portRepository.Insert(
        code,
        name,
        Type,
        city_Id,
        state,
        latitude,
        longitude,
    )
    log.Println(Code)
    if err != nil {
        return services.Port{}, err
    }
    return s.portRepoToService(
        repoPort,
    ), nil
}
func (s *service) Get(code string) (services.Port, error) {
    repoPort, exists, err := s.portRepository.ReadOne(code)
    if err != nil {
        return services.Port{}, err
    }
    if !exists {
        return services.Port{}, fmt.Errorf(
            "no port found for code: %s",
            code,
        )
    }
    return s.portRepoToService(
        repoPort,
    ), nil
}
func (s *service) Modify(
    code string,
    Name *string,
    Type *string,
    City_Id *int64,
    State *string,
    Latitude *float64,
    Longitude *float64,
) (services.Port, error) {
    if len(code) != 0 {
        _, exists, err := s.portRepository.ReadOne(code)
        if err != nil {
            return services.Port{}, err
        }
        if !exists {
            return services.Port{}, fmt.Errorf(
                "could not find port with code: %s",
                code,
            )
        }
    }
    _, err := s.portRepository.Update(
        code,
        Name,
        Type,
        City_Id,
        State,
        Latitude,
        Longitude,
    )
    if err != nil {
        return services.Port{}, err
    }
    return s.Get(code)
}
func (s *service) GetMany(
        pageNumber *uint,
        itemsPerPage uint,
    ) ([]services.Port, error) {
    repoPorts , err := s.portRepository.ReadMany(
        pageNumber,
        itemsPerPage,
    )
    if err != nil {
        return make([]services.Port,0),err
    }
    ports := []services.Port{}
    for _, cr := range repoPorts {
        ports = append(
            ports,
            s.portRepoToService(
                cr,
            ),
        )
    }
    return ports, nil
}
func (s *service) Remove(code string) error {
    return s.portRepository.Delete(code)
}
func (s *service) portRepoToService(
    repoPort repositories.Ports,
) services.Port {
    return services.Port{
        Code: repoPort.Code,
        Name: repoPort.Name,
        Type: repoPort.Type,
        City_Id: repoPort.City_Id,
        State: repoPort.State,
        Latitude: repoPort.Latitude,
        Longitude: repoPort.Longitude,
    }
}