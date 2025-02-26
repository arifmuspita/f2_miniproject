package model

import (
	"time"
)

type WalletHistory struct {
	ID              int     `gorm:"primaryKey"`
	UserID          int     `gorm:"not null;index" json:"user_id"`
	User            User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Amount          float64 `gorm:"not null" json:"amount"`
	TransactionType string  `gorm:"type:varchar(10);not null" json:"transaction_type"` // "income" or "expense"
	SourceType      string  `gorm:"type:varchar(10);not null" json:"source_type"`      // "rental", "rental_income", "deposit", "withdraw"
	SourceID        int     `gorm:"not null" json:"source_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
