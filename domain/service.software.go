package domain

import "sort"

type SoftwareListDTO struct {
	Software
	LatestRelease *Release
}

type SoftwareServiceInterface interface {
	List() ([]SoftwareListDTO, error)
	GetByName(name string) (*Software, error)
	GetLatestRelease(softwareID string) (*Release, error)
}

type SoftwareService struct {
	softeare_repository SoftwareRepositoryInterface
}

func NewSoftwareService(softeare_repository SoftwareRepositoryInterface) *SoftwareService {
	return &SoftwareService{softeare_repository: softeare_repository}
}

func (s *SoftwareService) List() ([]SoftwareListDTO, error) {
	softwares, err := s.softeare_repository.List()
	if err != nil {
		return nil, err
	}

	dtos := make([]SoftwareListDTO, 0, len(softwares))
	for _, sw := range softwares {
		dto := SoftwareListDTO{Software: sw}
		if len(sw.Releases) > 0 {
			sort.Slice(sw.Releases, func(i, j int) bool {
				return sw.Releases[i].CreatedAt.After(sw.Releases[j].CreatedAt)
			})
			dto.LatestRelease = &sw.Releases[0]
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (s *SoftwareService) GetByName(name string) (*Software, error) {
	return s.softeare_repository.GetByName(name)
}

func (s *SoftwareService) GetLatestRelease(softwareID string) (*Release, error) {
	software, err := s.softeare_repository.GetByName(softwareID)
	if err != nil {
		return nil, err
	}
	if software == nil {
		return nil, nil
	}
	if len(software.Releases) == 0 {
		return nil, nil
	}
	sort.Slice(software.Releases, func(i, j int) bool {
		return software.Releases[i].CreatedAt.After(software.Releases[j].CreatedAt)
	})
	return &software.Releases[0], nil
}
