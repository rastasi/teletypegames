package domain

import "gorm.io/gorm"

type Software struct {
	gorm.Model
	Name          string         `gorm:"uniqueIndex;size:128" json:"name"`
	Title         string         `gorm:"size:255" json:"title"`
	Author        string         `gorm:"size:255" json:"author"`
	Desc          string         `gorm:"type:text" json:"desc"`
	Story         string         `gorm:"type:text" json:"story"`
	License       string         `gorm:"size:128" json:"license"`
	Platform      string         `gorm:"size:128" json:"platform"`
	Status        string         `gorm:"size:20;default:development" json:"status"`
	Highlighted   bool           `gorm:"default:false" json:"highlighted"`
	Releases      []Release      `json:"-"`
	ExternalLinks []ExternalLink `json:"externalLinks"`
}
