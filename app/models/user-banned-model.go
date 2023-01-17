package models

import "time"

type UserBanned struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint
	User        User
	ModeratorId uint      `gorm:"default:0;index"`
	UpdatedAt   time.Time `gorm:"type:DATETIME;default:NULL"`
	CreatedAt   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	IsDeleted   bool      `gorm:"default:false;index;index:idx_bool"`
	DeletedAt   time.Time `gorm:"default:NULL"`
}
