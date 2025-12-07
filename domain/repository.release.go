package domain

import "gorm.io/gorm"

type ReleaseRepositoryInterface interface {
	Create(release *Release) error
	FindLatestBySoftwareID(softwareID uint) (*Release, error)
	FindBySoftwareIDAndVersion(softwareID uint, version string) (*Release, error)
}

type ReleaseRepository struct {
	db *gorm.DB
}

func (r *ReleaseRepository) Create(release *Release) error {
	return r.db.Create(release).Error
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
