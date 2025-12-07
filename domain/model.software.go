package domain

import "gorm.io/gorm"

type Software struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex;size:128" json:"name"`
	Title    string    `gorm:"size:255" json:"title"`
	Author   string    `gorm:"size:255" json:"author"`
	Desc     string    `gorm:"type:text" json:"desc"`
	Site     string    `gorm:"size:255" json:"site"`
	License  string    `gorm:"size:128" json:"license"`
	Platform string    `gorm:"size:128" json:"platform"`
	Releases []Release `json:"releases"`
}
