package database

import (
	"log"
	"time"

	"github.com/Carmo-sousa/webAPI/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dns := "host=localhost port=5432 user=admin dbname=books sslmode=disable password=199718"

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxIdleTime(time.Hour)

	migrations.RunMigrate(db)
}

func CloseConn() error {
	config, err := db.DB()

	if err != nil {
		return err
	}
	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDataBase() *gorm.DB {
	return db
}
