package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobStatus string

const (
	JobQueued     JobStatus = "QUEUED"
	JobProcessing JobStatus = "PROCESSING"
	JobCompleted  JobStatus = "COMPLETED"
	JobFailed     JobStatus = "FAILED"
)

type Job struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Status    JobStatus `gorm:"type:varchar(20);" json:"status"`
	Filename  string    `gorm:"type:varchar(255);" json:"filename"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (job *Job) BeforeCreate(tx *gorm.DB) (err error) {
	for {
		job.ID = uuid.New()
		var count int64
		if err := tx.Model(&Job{}).Where("id = ?", job.ID).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			break
		}
	}
	job.CreatedAt = time.Now()
	return nil
}

func (job *Job) BeforeUpdate(tx *gorm.DB) (err error) {
	job.UpdatedAt = time.Now()
	return nil
}
