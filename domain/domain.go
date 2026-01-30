package domain

import (
	"teletype_softwares/lib/mysql_utils"
)

type Domain struct {
	SoftwareRepository     SoftwareRepositoryInterface
	ReleaseRepository      ReleaseRepositoryInterface
	FileRepository         FileRepositoryInterface
	DownloadService        DownloadServiceInterface
	SoftwareUpdaterService SoftwareUpdaterServiceInterface
	SoftwareService        SoftwareServiceInterface
	FileService            FileServiceInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateGoDatabase(DB)

	software_repository := &SoftwareRepository{db: DB}
	release_repository := &ReleaseRepository{db: DB}
	file_repository := NewFileRepository()

	software_service := NewSoftwareService(software_repository, release_repository)

	tic80_updater := NewSoftwareUpdaterTIC80Service(
		software_repository,
		release_repository,
		file_repository,
	)

	ebitengine_updater := NewSoftwareUpdaterEbitengineService(
		software_repository,
		release_repository,
		file_repository,
	)
	software_updater_service := NewSoftwareUpdaterService(tic80_updater, ebitengine_updater)

	download_service := NewDownloadService(software_repository, release_repository)

	file_service := NewFileService(file_repository)

	return Domain{
		SoftwareRepository:     software_repository,
		ReleaseRepository:      release_repository,
		FileRepository:         file_repository,
		DownloadService:        download_service,
		SoftwareUpdaterService: software_updater_service,
		SoftwareService:        software_service,
		FileService:            file_service,
	}
}
