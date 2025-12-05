package domain

import (
	"teletype_softwares/lib/mysql_utils"
)

type Domain struct {
	SoftwareRepository     SoftwareRepository
	ReleaseRepository      ReleaseRepository
	DownloadService        DownloadService
	SoftwareUpdaterService SoftwareUpdaterService
	SoftwareService        SoftwareService
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateGoDatabase(DB)

	softwareRepository := &softwareRepository{db: DB}
	releaseRepository := &releaseRepository{db: DB}
	softwareService := &softwareService{softwareRepository: softwareRepository}
	tic80Updater := &softwareUpdaterTIC80Service{
		softwareRepository: softwareRepository,
		releaseRepository:  releaseRepository,
	}

	softwareUpdaterService := &softwareUpdaterService{tic80Updater: tic80Updater}

	downloadServiceInstance := &downloadService{
		softwareRepository: softwareRepository,
		releaseRepository:  releaseRepository,
	}

	return Domain{
		SoftwareRepository:     softwareRepository,
		ReleaseRepository:      releaseRepository,
		DownloadService:        downloadServiceInstance,
		SoftwareUpdaterService: softwareUpdaterService,
		SoftwareService:        softwareService,
	}
}
