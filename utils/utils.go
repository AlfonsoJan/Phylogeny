package utils

import (
	"Phylogeny/entities/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUniqueUUID(db *gorm.DB) (uuid.UUID, error) {
	for {
		newID := uuid.New()
		var count int64
		if err := db.Model(&models.Job{}).Where("id = ?", newID).Count(&count).Error; err != nil {
			return uuid.Nil, err
		}
		if count == 0 {
			return newID, nil
		}
	}
}
