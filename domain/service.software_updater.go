package domain

import (
	"fmt"
)

type SoftwareUpdaterServiceInterface interface {
	Update(platform, name, version string) error
}

type SoftwareUpdaterService struct {
	tic80Updater SoftwareUpdaterTIC80ServiceInterface
}

func NewSoftwareUpdaterService(tic80Updater SoftwareUpdaterTIC80ServiceInterface) *SoftwareUpdaterService {
	return &SoftwareUpdaterService{tic80Updater: tic80Updater}
}

func (s *SoftwareUpdaterService) Update(platform, name, version string) error {
	if platform == "tic80" {
		return s.tic80Updater.Update(name, version)
	}
	return fmt.Errorf("unsupported platform: %s", platform)
}
