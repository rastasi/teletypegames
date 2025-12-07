package domain

import (
	"teletype_softwares/lib/mysql_utils"
)

type Domain struct {
	SoftwareRepository            SoftwareRepositoryInterface
	ReleaseRepository             ReleaseRepositoryInterface
	FileRepository                FileRepositoryInterface
	DownloadService               DownloadServiceInterface
	SoftwareUpdaterService        SoftwareUpdaterServiceInterface
	SoftwareService               SoftwareServiceInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateGoDatabase(DB)

	softwareRepository := &SoftwareRepository{db: DB}
	releaseRepository := &ReleaseRepository{db: DB}
	fileRepository := NewFileRepository()

	softwareService := NewSoftwareService(softwareRepository)
	
	tic80Updater := NewSoftwareUpdaterTIC80Service(
		softwareRepository,
		releaseRepository,
		fileRepository,
	)

	softwareUpdaterService := NewSoftwareUpdaterService(tic80Updater)

	downloadService := NewDownloadService(softwareRepository, releaseRepository)

	return Domain{
		SoftwareRepository:     softwareRepository,
		ReleaseRepository:      releaseRepository,
		FileRepository:         fileRepository,
		DownloadService:        downloadService,
		SoftwareUpdaterService: softwareUpdaterService,
		SoftwareService:        softwareService,
	}
}
