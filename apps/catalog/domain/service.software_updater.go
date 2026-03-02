package domain

import (
	"fmt"
)

type SoftwareUpdaterServiceInterface interface {
	Update(platform, name, version string) error
}

type SoftwareUpdaterService struct {
	tic80Updater      SoftwareUpdaterTIC80ServiceInterface
	ebitengineUpdater SoftwareUpdaterEbitengineServiceInterface
	loveUpdater       SoftwareUpdaterLoveServiceInterface
}

func NewSoftwareUpdaterService(tic80_updater SoftwareUpdaterTIC80ServiceInterface, ebitengine_updater SoftwareUpdaterEbitengineServiceInterface, love_updater SoftwareUpdaterLoveServiceInterface) *SoftwareUpdaterService {
	return &SoftwareUpdaterService{
		tic80Updater:      tic80_updater,
		ebitengineUpdater: ebitengine_updater,
		loveUpdater:       love_updater,
	}
}

func (s *SoftwareUpdaterService) Update(platform, name, version string) error {
	if platform == "tic80" {
		return s.tic80Updater.Update(name, version)
	}
	if platform == "ebitengine" {
		return s.ebitengineUpdater.Update(name, version)
	}
	if platform == "love" {
		return s.loveUpdater.Update(name, version)
	}
	return fmt.Errorf("unsupported platform: %s", platform)
}
