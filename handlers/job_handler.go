package handlers

import (
	"Phylogeny/database/queries"
	"Phylogeny/models"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateJobHandler creates a new job with an uploaded file
//
//	@Description	Upload a file to create a new job with its filename stored in the database
//	@Summary		create a new job
//	@Tags			job
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			file	formData	file					true	"File to be uploaded"
//	@Success		201		{object}	models.Job				"Job successfully created"
//	@Failure		400		{object}	models.ErrorResponse	"Bad request, file missing in the form-data"
//	@Failure		500		{object}	models.ErrorResponse	"Internal server error, unable to create job"
//	@Router			/job [post]
func CreateJobHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Failed to get file",
			Message: err.Error(),
		})
	}

	tempDir := filepath.Join(os.TempDir(), "phylogeny")

	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Cannot create temp directory",
			Message: err.Error(),
		})
	}

	fileUUID := uuid.New()
	newFilename := fmt.Sprintf("%s_%s", fileUUID.String(), fileHeader.Filename)
	filePath := filepath.Join(tempDir, newFilename)

	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to save file",
			Message: err.Error(),
		})
	}

	job := new(models.Job)
	job.Filename = newFilename
	job.Status = models.JobQueued

	if err := queries.CreateJob(job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Cannot create job",
			Message: err.Error(),
		})
	}

	log.Println("Job created with ID:", job.ID)
	return c.Status(fiber.StatusCreated).JSON(job)
}
