package domain

import "gorm.io/gorm"

type ReleaseRepositoryInterface interface {
	Create(release *Release) error
	CreateIfNotExist(release *Release) error
	FindLatestBySoftwareID(softwareID uint) (*Release, error)
	FindBySoftwareIDAndVersion(softwareID uint, version string) (*Release, error)
	ListBySoftwareID(softwareID uint) []Release
}

type ReleaseRepository struct {
	db *gorm.DB
}

func (r *ReleaseRepository) Create(release *Release) error {
	return r.db.Create(release).Error
}

func (r *ReleaseRepository) CreateIfNotExist(release *Release) error {
	var existing Release
	err := r.db.Where("software_id = ? AND version = ?", release.SoftwareID, release.Version).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Create(release).Error
		}
		return err
	}
	return nil
}

func (r *ReleaseRepository) FindLatestBySoftwareID(softwareID uint) (*Release, error) {
	var release Release
	err := r.db.Where("software_id = ?", softwareID).Order("created_at desc").First(&release).Error
	return &release, err
}

func (r *ReleaseRepository) FindBySoftwareIDAndVersion(softwareID uint, version string) (*Release, error) {
	var release Release
	err := r.db.Where("software_id = ? AND version = ?", softwareID, version).First(&release).Error
	return &release, err
}

func (r *ReleaseRepository) ListBySoftwareID(softwareID uint) []Release {
	var releases []Release
	r.db.Where("software_id = ?", softwareID).Find(&releases)
	return releases
}
