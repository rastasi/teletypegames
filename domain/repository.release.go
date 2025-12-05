package domain

import "gorm.io/gorm"

type ReleaseRepository interface {
	Create(release *Release) error
	FindLatestBySoftwareID(softwareID uint) (*Release, error)
	FindBySoftwareIDAndVersion(softwareID uint, version string) (*Release, error)
}

type releaseRepository struct {
	db *gorm.DB
}

// NewReleaseRepository creates a new instance of ReleaseRepository.
/* func NewReleaseRepository(db *gorm.DB) ReleaseRepository {
	return &releaseRepository{db}
} */

func (r *releaseRepository) Create(release *Release) error {
	return r.db.Create(release).Error
}

func (r *releaseRepository) FindLatestBySoftwareID(softwareID uint) (*Release, error) {
	var release Release
	err := r.db.Where("software_id = ?", softwareID).Order("created_at desc").First(&release).Error
	return &release, err
}

func (r *releaseRepository) FindBySoftwareIDAndVersion(softwareID uint, version string) (*Release, error) {
	var release Release
	err := r.db.Where("software_id = ? AND version = ?", softwareID, version).First(&release).Error
	return &release, err
}
