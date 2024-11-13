package handlers

import (
	"Phylogeny/database"
	"Phylogeny/database/queries"
	"Phylogeny/entities/dto"
	"Phylogeny/entities/models"
	"Phylogeny/tasks"
	"Phylogeny/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

var JobQueue *tasks.JobQueue

// CreateJobHandler handles the creation of a new job with a file upload.
//
// @Summary Create a new job with file upload
// @Description This endpoint allows you to create a new job by uploading a file, which will be saved in a temporary directory.
// @Tags jobs
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "File to upload"
// @Success 201 {object} models.Job "Job created successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad request, file missing in the form-data"
// @Failure 500 {object} dto.ErrorResponse "Internal server error, unable to create job"
// @Router /jobs [post]
func CreateJobHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "Failed to get file",
			Message: err.Error(),
		})
	}
	tempDir := filepath.Join(os.TempDir(), "phylogeny")

	job := new(models.Job)
	job.Status = models.JobQueued
	job.ID, err = utils.GenerateUniqueUUID(database.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "Cannot generate UUID",
			Message: err.Error(),
		})
	}

	newFilename := fmt.Sprintf("%s_%s", job.ID.String(), fileHeader.Filename)
	filePath := filepath.Join(tempDir, newFilename)
	job.Filename = newFilename

	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "Failed to save file",
			Message: err.Error(),
		})
	}

	if err := queries.CreateJob(job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "Cannot create job",
			Message: err.Error(),
		})
	}

	JobQueue.Enqueue(job)

	log.Println("Job created with ID:", job.ID)
	return c.Status(fiber.StatusCreated).JSON(job)
}
