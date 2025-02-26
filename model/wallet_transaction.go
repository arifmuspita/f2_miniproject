package model

import (
	"time"
)

type WalletTransaction struct {
	ID              int     `gorm:"primaryKey"`
	UserID          int     `gorm:"not null;index" json:"user_id"`
	User            User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TransactionType string  `gorm:"type:varchar(10);not null" json:"transaction_type"` // "deposit" or "withdraw"
	Amount          float64 `gorm:"not null" json:"amount"`
	Status          string  `gorm:"type:varchar(10);default:'pending'" json:"status"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
