package domain

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	SoftwareID     uint   `gorm:"index" json:"software_id"`
	Version        string `gorm:"size:64" json:"version"`
	CartridgePath  string `gorm:"size:255" json:"-"`
	SourcePath     string `gorm:"size:255" json:"-"`
	HTMLFolderPath string `gorm:"size:255" json:"-"`
	DocsFolderPath string `gorm:"size:255" json:"-"`
}
