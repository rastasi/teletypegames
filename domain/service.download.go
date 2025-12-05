package domain

import (
	"errors"
)

var (
	ErrSoftwareNotFound = errors.New("software not found")
	ErrNoReleasesFound  = errors.New("no releases found for software")
	ErrReleaseNotFound  = errors.New("release not found for software and version")
)

type DownloadService interface {
	GetLatestRelease(softwareName string) (*Release, error)
	GetSpecificRelease(softwareName string, version string) (*Release, error)
}

type downloadService struct {
	softwareRepository SoftwareRepository
	releaseRepository  ReleaseRepository
}

func (s *downloadService) GetLatestRelease(softwareName string) (*Release, error) {
	software, err := s.softwareRepository.GetByName(softwareName)
	if err != nil {
		return nil, ErrSoftwareNotFound
	}

	release, err := s.releaseRepository.FindLatestBySoftwareID(software.ID)
	if err != nil {
		return nil, ErrNoReleasesFound
	}

	return release, nil
}

func (s *downloadService) GetSpecificRelease(softwareName string, version string) (*Release, error) {
	software, err := s.softwareRepository.GetByName(softwareName)
	if err != nil {
		return nil, ErrSoftwareNotFound
	}

	release, err := s.releaseRepository.FindBySoftwareIDAndVersion(software.ID, version)
	if err != nil {
		return nil, ErrReleaseNotFound
	}

	return release, nil
}
