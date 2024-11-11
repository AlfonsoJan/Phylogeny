package queries

import (
	"Phylogeny/database"
	"Phylogeny/models"
	"Phylogeny/utils"
	"errors"
	"time"
)

func CreateJob(job *models.Job) error {
	job.CreatedAt = time.Now()
	job.Status = models.JobQueued
	return database.DB.Create(&job).Error
}

func GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	if err := database.DB.First(&job, "id = ?", id).Error; err != nil {
		return nil, errors.New("job not found")
	}
	return &job, nil
}

func UpdateJob(id string, jobData *models.Job) (*models.Job, error) {
	var job models.Job
	if err := database.DB.First(&job, "id = ?", id).Error; err != nil {
		return nil, errors.New("job not found")
	}
	job.Status = utils.If(jobData.Status != "", jobData.Status, job.Status)
	job.Filename = utils.If(jobData.Filename != "", jobData.Filename, job.Filename)
	job.UpdatedAt = time.Now()
	if err := database.DB.Save(&job).Error; err != nil {
		return nil, errors.New("could not update job")
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
