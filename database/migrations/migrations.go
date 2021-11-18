package migrations

import (
	"log"

	"github.com/Carmo-sousa/webAPI/models"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{})

	if err != nil {
		log.Fatal(err)
	}
}
