package model

import (
	"time"
)

type TransactionDetail struct {
	ID            int         `gorm:"primaryKey"`
	TransactionID int         `gorm:"not null;index" json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	ToolID        int         `gorm:"not null;index" json:"tool_id"`
	Tool          Tool        `gorm:"foreignKey:ToolID" json:"tool,omitempty"`
	Quantity      int         `gorm:"not null" json:"quantity"`
	Subtotal      float64     `gorm:"not null" json:"sub_total"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
