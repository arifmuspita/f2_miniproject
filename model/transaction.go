package model

import (
	"time"
)

type Transaction struct {
	ID         int       `gorm:"primaryKey"`
	RenterID   int       `gorm:"not null;index" json:"renter_id"`
	StartDate  time.Time `gorm:"not null" json:"start_date"`
	EndDate    time.Time `gorm:"not null" json:"end_date"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	Status     string    `gorm:"type:varchar(10);default:'pending'" json:"status"`
	Renter     User      `gorm:"foreignKey:RenterID" json:"owner,omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
