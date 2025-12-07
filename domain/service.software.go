package domain

import "sort"

type SoftwareListDTO struct {
	Software
	LatestRelease *Release
}

type SoftwareServiceInterface interface {
	List() ([]SoftwareListDTO, error)
	GetByNameWithReleases(name string) (*Software, error)
}

type SoftwareService struct {
	repository SoftwareRepositoryInterface
}

func NewSoftwareService(repository SoftwareRepositoryInterface) *SoftwareService {
	return &SoftwareService{repository: repository}
}

func (s *SoftwareService) List() ([]SoftwareListDTO, error) {
	softwares, err := s.repository.List()
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

func (s *SoftwareService) GetByNameWithReleases(name string) (*Software, error) {
	return s.repository.GetByName(name)
}
