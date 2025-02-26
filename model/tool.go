package model

import (
	"time"

	"gorm.io/gorm"
)

type Tool struct {
	ID           int      `gorm:"primaryKey"`
	Name         string   `gorm:"type:varchar(100);not null" json:"name"`
	Price        float64  `gorm:"not null" json:"price"`
	Availability bool     `gorm:"default:true" json:"available"`
	Stock        int      `gorm:"not null;default:1" json:"stock"`
	CategoryID   int      `gorm:"not null;index" json:"categori_id"`
	Category     Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	OwnerID      int      `gorm:"not null;index" json:"owner_id"`
	Owner        User     `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
