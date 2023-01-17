package models

import "time"

type User struct {
	Id          uint `gorm:"primaryKey"`
	GroupId     uint
	Group       Group
	Username    string    `gorm:"type:VARCHAR(100);not null;index:,class:FULLTEXT"`
	Password    string    `gorm:"type:VARCHAR(100);not null"`
	Email       string    `gorm:"type:VARCHAR(150);not null;index:,unique"`
	Name        string    `gorm:"type:VARCHAR(50);default:NULL"`
	Surname     string    `gorm:"type:VARCHAR(50);default:NULL"`
	Age         uint8     `gorm:"type:TINYINT(2);default:1"`
	Gender      string    `gorm:"type:VARCHAR(10);default:'none'"`
	IsConfirm   *bool     `gorm:"default:false;index;index:idx_bool"`
	ModeratorId uint      `gorm:"default:0;index"`
	IsArchive   *bool     `gorm:"default:false;index;index:idx_bool"`
	ArchivedAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	IsBlocked   *bool     `gorm:"default:false;index;index:idx_bool"`
	ActivityAt  time.Time `gorm:"type:DATETIME;default:NULL"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   *bool     `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
}
