package queries

import (
	"Phylogeny/database"
	"Phylogeny/models"
	"errors"
)

func CreateJob(job *models.Job) error {
	job.Status = models.JobQueued
	return database.DB.Create(&job).Error
}

func updateJobStatus(job *models.Job, status models.JobStatus) error {
	job.Status = status
	return database.DB.Save(&job).Error
}

func GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	if err := database.DB.First(&job, "id = ?", id).Error; err != nil {
		return nil, errors.New("job not found")
	}
	return &job, nil
}

func DeleteJob(id string) error {
	var job models.Job
	if err := database.DB.First(&job, "id = ?", id).Error; err != nil {
		return errors.New("job not found")
	}
	return database.DB.Delete(&job).Error
}
