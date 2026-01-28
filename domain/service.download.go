package domain

import (
	"errors"
)

var (
	ErrSoftwareNotFound = errors.New("software not found")
	ErrNoReleasesFound  = errors.New("no releases found for software")
	ErrReleaseNotFound  = errors.New("release not found for software and version")
)

type DownloadServiceInterface interface {
	GetLatestRelease(software_name string) (*Release, error)
	GetSpecificRelease(software_name string, version string) (*Release, error)
}

type DownloadService struct {
	softwareRepository SoftwareRepositoryInterface
	releaseRepository  ReleaseRepositoryInterface
}

func NewDownloadService(software_repository SoftwareRepositoryInterface, release_repository ReleaseRepositoryInterface) *DownloadService {
	return &DownloadService{
		softwareRepository: software_repository,
		releaseRepository:  release_repository,
	}
}

func (s *DownloadService) GetLatestRelease(software_name string) (*Release, error) {
	software, err := s.softwareRepository.GetByName(software_name)
	if err != nil {
		return nil, ErrSoftwareNotFound
	}

	release, err := s.releaseRepository.FindLatestBySoftwareID(software.ID)
	if err != nil {
		return nil, ErrNoReleasesFound
	}

	return release, nil
}

func (s *DownloadService) GetSpecificRelease(software_name string, version string) (*Release, error) {
	software, err := s.softwareRepository.GetByName(software_name)
	if err != nil {
		return nil, ErrSoftwareNotFound
	}

	release, err := s.releaseRepository.FindBySoftwareIDAndVersion(software.ID, version)
	if err != nil {
		return nil, ErrReleaseNotFound
	}

	return release, nil
}
