package model

import (
	"time"
)

type Notification struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"not null;index" json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Message   string `gorm:"type:varchar(255);not null" json:"message"`
	Status    string `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
