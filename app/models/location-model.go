package models

import "time"

type Location struct {
	Id          uint `gorm:"primaryKey"`
	LanguageId  uint
	Language    Language
	UserId      uint
	User        User
	CountryId   uint
	Country     Country
	Name        string    `gorm:"type:VARCHAR(255);not null;index:,class:FULLTEXT"`
	City        string    `gorm:"type:VARCHAR(150);not null;index:,class:FULLTEXT"`
	Description string    `gorm:"type:TEXT;default:null"`
	IsVisible   *bool     `gorm:"default:true;index;index:idx_bool"`
	IsArchive   *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"default:NULL"`
}
