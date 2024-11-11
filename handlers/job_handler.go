package handlers

import (
	"Phylogeny/database/queries"
	"Phylogeny/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateJobHandler creates a new job
// @Description Create a new job
// @Summary create a new job
// @Tags job
// @Accept json
// @Produce json
// @Param job body models.Job true "Job object that needs to be created"
// @Success 201 {object} models.Job
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /job [post]
func CreateJobHandler(c *fiber.Ctx) error {
	job := new(models.Job)
	if err := c.BodyParser(job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	job.Status = models.JobQueued
	if err := queries.CreateJob(job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create job"})
	}
	log.Println("Job created with ID:", job.ID)
	return c.Status(fiber.StatusCreated).JSON(job)
}

// GetJobHandler retrieves a job by its ID
// @Description Get a job by its ID
// @Summary retrieve job by ID
// @Tags job
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} models.Job
// @Failure 404 {object} error
// @Router /job/{id} [get]
func GetJobHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := queries.GetJobByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("Job found with ID:", job.ID)
	return c.JSON(job)
}

// UpdateJobHandler updates an existing job by its ID
// @Description Update a job by its ID
// @Summary update job by ID
// @Tags job
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Param job body models.Job true "Job object with updated fields"
// @Success 200 {object} models.Job
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /job/{id} [put]
func UpdateJobHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var jobUpdate models.Job

	if err := c.BodyParser(&jobUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	updatedJob, err := queries.UpdateJob(id, &jobUpdate)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("Job updated with ID:", updatedJob.ID)
	return c.JSON(updatedJob)
}

// DeleteJobHandler deletes a job by its ID
// @Description Delete a job by its ID
// @Summary delete job by ID
// @Tags job
// @Param id path string true "Job ID"
// @Success 204 "No Content"
// @Failure 404 {object} error
// @Router /job/{id} [delete]
func DeleteJobHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := queries.DeleteJob(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("Job deleted with ID:", id)
	return c.SendStatus(fiber.StatusNoContent)
}
