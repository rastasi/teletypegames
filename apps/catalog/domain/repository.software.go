package domain

import "gorm.io/gorm"

type SoftwareRepositoryInterface interface {
	List() ([]Software, error)
	GetByName(name string) (*Software, error)
	UpdateOrCreate(software *Software) error
}

type SoftwareRepository struct {
	db *gorm.DB
}

func (r *SoftwareRepository) GetByName(name string) (*Software, error) {
	var software Software
	if err := r.db.Preload("Releases").Where("name = ?", name).First(&software).Error; err != nil {
		return nil, err
	}
	return &software, nil
}

func (r *SoftwareRepository) List() ([]Software, error) {
	var softwares []Software
	if err := r.db.Preload("Releases").Find(&softwares).Error; err != nil {
		return nil, err
	}
	return softwares, nil
}

func (r *SoftwareRepository) UpdateOrCreate(software *Software) error {
	var existing Software
	if err := r.db.Where("name = ?", software.Name).First(&existing).Error; err == nil {
		software.ID = existing.ID
		return r.db.Model(&existing).Updates(software).Error
	} else {
		return r.db.Create(software).Error
	}
}
