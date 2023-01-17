package models

import "time"

type Auth struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint
	User      User
	Ip        string    `gorm:"type:VARCHAR(100);default:null;index"`
	Device    string    `gorm:"type:VARCHAR(30);default:null;index"`
	UserAgent string    `gorm:"type:TINYTEXT;default:null;index:,class:FULLTEXT"`
	CreatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
}
