package domain

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey" json:"id"`
	SoftwareID    uint   `gorm:"index" json:"software_id"`
	Version       string `gorm:"size:64" json:"version"`
	CartridgePath string `gorm:"size:255" json:"-"`
	SourcePath    string `gorm:"size:255" json:"-"`
}
