package domain

import "sort"

type SoftwareListDTO struct {
	Software
	LatestRelease *Release
}

type SoftwareService interface {
	List() ([]SoftwareListDTO, error)
	GetByNameWithReleases(name string) (*Software, error)
}

type softwareService struct {
	softwareRepository SoftwareRepository
}

func (s *softwareService) List() ([]SoftwareListDTO, error) {
	softwares, err := s.softwareRepository.List()
	if err != nil {
		return nil, err
	}

	dtos := make([]SoftwareListDTO, 0, len(softwares))
	for _, sw := range softwares {
		dto := SoftwareListDTO{Software: sw}
		if len(sw.Releases) > 0 {
			// Sort releases by CreatedAt in descending order to find the latest
			sort.Slice(sw.Releases, func(i, j int) bool {
				return sw.Releases[i].CreatedAt.After(sw.Releases[j].CreatedAt)
			})
			dto.LatestRelease = &sw.Releases[0]
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (s *softwareService) GetByNameWithReleases(name string) (*Software, error) {
	return s.softwareRepository.GetByName(name)
}
