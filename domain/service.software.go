package domain

import (
	"fmt"
	"sort"
)

type SoftwareListDTO struct {
	Software
	LatestRelease *Release
}

type SoftwareDetaildListDTO struct {
	Softwares []SoftwareShowData `json:"softwares"`
}

type SoftwareShowData struct {
	Software           *Software `json:"software"`
	Releases           []Release `json:"releases"`
	LatestRelease      *Release  `json:"latestRelease"`
	WebPlayableRelease *Release  `json:"webPlayableRelease"`
}

type SoftwareServiceInterface interface {
	List() ([]SoftwareListDTO, error)
	DetailedList() (*SoftwareDetaildListDTO, error)
	GetByName(name string) (*Software, error)
	GetLatestRelease(software_id string) (*Release, error)
	GetForShowByName(name string) (*SoftwareShowData, error)
}

type SoftwareService struct {
	softeare_repository SoftwareRepositoryInterface
	release_repository  ReleaseRepositoryInterface
}

func NewSoftwareService(
	softeare_repository SoftwareRepositoryInterface,
	release_repository ReleaseRepositoryInterface,
) *SoftwareService {
	return &SoftwareService{
		softeare_repository: softeare_repository,
		release_repository:  release_repository,
	}
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

func (s *SoftwareService) GetForShowByName(name string) (*SoftwareShowData, error) {
	software, err := s.softeare_repository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if software == nil {
		return nil, fmt.Errorf("software not found")
	}

	releases := s.release_repository.ListBySoftwareID(software.ID)

	var latest_release *Release
	if len(releases) > 0 {
		sort.Slice(releases, func(i, j int) bool {
			return releases[i].CreatedAt.After(releases[j].CreatedAt)
		})
		latest_release = &releases[0]
	}

	var web_payable_release *Release
	for _, release := range releases {
		if release.HTMLFolderPath != "" {
			web_payable_release = &release
			break
		}
	}

	return &SoftwareShowData{
		Software:           software,
		Releases:           releases,
		LatestRelease:      latest_release,
		WebPlayableRelease: web_payable_release,
	}, nil
}

func (s *SoftwareService) GetLatestRelease(software_id string) (*Release, error) {
	software, err := s.softeare_repository.GetByName(software_id)
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

func (s *SoftwareService) DetailedList() (*SoftwareDetaildListDTO, error) {
	softwares, err := s.softeare_repository.List()
	if err != nil {
		return nil, err
	}

	dtos := make([]SoftwareShowData, 0, len(softwares))
	for _, software := range softwares {
		showData, err := s.GetForShowByName(software.Name)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, *showData)
	}

	return &SoftwareDetaildListDTO{Softwares: dtos}, nil
}
