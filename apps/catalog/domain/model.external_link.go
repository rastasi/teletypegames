package domain

import "gorm.io/gorm"

type ExternalLink struct {
	gorm.Model
	SoftwareID uint   `gorm:"index" json:"softwareId"`
	Label      string `gorm:"size:128" json:"label"`
	URL        string `gorm:"size:255" json:"url"`
}
