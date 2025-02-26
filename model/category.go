package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique" json:"name"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
