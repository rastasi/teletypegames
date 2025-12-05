package domain

import (
	"fmt"
)

// Keep os for getting GAME_PATH in TIC80 service if needed

type SoftwareUpdaterService interface {
	UpdateSoftware(platform, name string) error
}

type softwareUpdaterService struct {
	tic80Updater SoftwareUpdaterTIC80Service
}

func (s *softwareUpdaterService) UpdateSoftware(platform, name string) error {
	if platform == "tic80" {
		return s.tic80Updater.UpdateTIC80Software(name)
	}
	return fmt.Errorf("unsupported platform: %s", platform)
}
