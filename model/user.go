package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int     `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(100);not null" json:"name"`
	Email        string  `gorm:"type:varchar(100);unique;not null" json:"email"`
	PasswordHash string  `gorm:"type:varchar(255);not null" json:"password"`
	Balance      float64 `gorm:"default:0" json:"balance"`
	IsVerified   bool    `gorm:"default:false" json:"is_verified"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
