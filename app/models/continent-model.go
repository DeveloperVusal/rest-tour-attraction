package models

import "time"

type Continent struct {
	Id          uint `gorm:"primaryKey"`
	LanguageId  uint
	Language    Language
	Name        string    `gorm:"type:VARCHAR(100);not null;index:,class:FULLTEXT"`
	ModeratorId uint      `gorm:"default:0;index"`
	IsVisible   *bool     `gorm:"default:true;index;index:idx_bool"`
	IsArchive   *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"default:NULL"`
}
