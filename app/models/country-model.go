package models

import "time"

type Country struct {
	Id          uint `gorm:"primaryKey"`
	ContinentId uint
	Continent   Continent
	LanguageId  uint
	Language    Language
	ModeratorId uint      `gorm:"default:0;index"`
	Name        string    `gorm:"type:VARCHAR(150);not null"`
	Capital     string    `gorm:"type:VARCHAR(150);null"`
	Code2       string    `gorm:"type:VARCHAR(2);not null;index"`
	Code3       string    `gorm:"type:VARCHAR(3);default:null;index"`
	Iso         uint      `gorm:"type:TINYINT(5);default:null;index"`
	IsVisible   *bool     `gorm:"default:true;index;index:idx_bool"`
	IsArchive   *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"default:NULL"`
}
