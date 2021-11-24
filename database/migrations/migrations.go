package migrations

import (
	"github.com/Carmo-sousa/webAPI/models"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
