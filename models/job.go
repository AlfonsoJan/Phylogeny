package models

import (
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
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Status   JobStatus `gorm:"type:varchar(20);" json:"status"`
	Filename string    `gorm:"type:varchar(255);" json:"filename"`
}
