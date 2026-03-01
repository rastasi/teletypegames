package domain

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	SoftwareID     uint   `gorm:"index" json:"softwareId"`
	Version        string `gorm:"size:64" json:"version"`
	CartridgePath  string `gorm:"size:255" json:"cartridgePath"`
	SourcePath     string `gorm:"size:255" json:"sourcePath"`
	HTMLFolderPath string `gorm:"size:255" json:"htmlFolderPath"`
	DocsFolderPath string `gorm:"size:255" json:"docsFolderPath"`
}
