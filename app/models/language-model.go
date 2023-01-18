package models

import "time"

type Language struct {
	Id          uint      `gorm:"primaryKey"`
	ModeratorId uint      `gorm:"default:0;index"`
	Name        string    `gorm:"type:VARCHAR(100);not null"`
	LangCode    string    `gorm:"type:VARCHAR(2);not null;index"`
	IsVisible   *bool     `gorm:"default:true;index;index:idx_bool"`
	IsArchive   *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"default:NULL"`
}
