package domain

import "gorm.io/gorm"

type ExternalLinkRepositoryInterface interface {
	UpsertByLabel(softwareID uint, label, url string) error
}

type ExternalLinkRepository struct {
	db *gorm.DB
}

func (r *ExternalLinkRepository) UpsertByLabel(softwareID uint, label, url string) error {
	var link ExternalLink
	err := r.db.Where("software_id = ? AND label = ?", softwareID, label).First(&link).Error
	if err == gorm.ErrRecordNotFound {
		return r.db.Create(&ExternalLink{SoftwareID: softwareID, Label: label, URL: url}).Error
	}
	if err != nil {
		return err
	}
	return r.db.Model(&link).Update("url", url).Error
}
