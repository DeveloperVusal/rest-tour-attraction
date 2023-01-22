package models

import "time"

type MainImage struct {
	Id               uint `gorm:"primaryKey"`
	LocationId       uint
	Location         Location
	UserId           uint
	User             User
	LanguageId       uint
	Language         Language
	IsMain           *bool     `gorm:"default:false;"`
	SourcePath       string    `gorm:"type:TEXT;not null"`
	ShortDescription string    `gorm:"type:VARCHAR(255);default:null"`
	IsVisible        *bool     `gorm:"default:true;index;index:idx_bool"`
	IsArchive        *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt       time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt        time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt        time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted        *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt        time.Time `gorm:"default:NULL"`
}
