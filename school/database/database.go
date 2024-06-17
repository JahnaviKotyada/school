package database

import (
	"log"
	"school/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=jahnavi@2003 dbname=school port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto Migrate the School model
	err = DB.AutoMigrate(&models.School{})
	if err != nil {
		log.Fatalf("failed to auto migrate models: %v", err)
	}

	return DB
}
