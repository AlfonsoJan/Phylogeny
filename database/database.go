package database

import (
	"Phylogeny/config"
	"Phylogeny/entities/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(postgres.Open(config.EnvConfig.DBURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	log.Println("Database connected!")
	if config.EnvConfig.DEV {
		dropTableIfDevelopment()
	}
	migrateColumns()
	return nil
}

func dropTableIfDevelopment() {
	if err := DB.Migrator().DropTable(&models.Job{}); err != nil {
		log.Fatalf("Failed to drop Job table: %v", err)
	}
}

func migrateColumns() {
	DB.AutoMigrate(&models.Job{})
}
