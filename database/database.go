package database

import (
	"Phylogeny/config"
	"Phylogeny/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	env := os.Getenv("ENV")
	dsn := config.GetDatabaseURL()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected!")

	if env == "dev" {
		DropTableIfDevelopment()
	}
	MigrateColumns()
}

func DropTableIfDevelopment() {
	if err := DB.Migrator().DropTable(&models.Job{}); err != nil {
		log.Fatalf("Failed to drop Job table: %v", err)
	}
}

func MigrateColumns() {
	DB.AutoMigrate(&models.Job{})
}
