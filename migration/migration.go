package migration

import (
	"f2_miniproject/model"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	if db == nil {
		log.Fatal("Database connection is nil! Migration aborted.")
		return
	}

	err := db.AutoMigrate(
		&model.Category{},
		&model.Notification{},
		&model.Tool{},
		&model.TransactionDetail{},
		&model.Transaction{},
		&model.User{},
		&model.WalletHistory{},
		&model.WalletTransaction{},
	)

	if err != nil {
		log.Fatal("Failed migration: ", err)
	}

	log.Println("Migration success!")
}
