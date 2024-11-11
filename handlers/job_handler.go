package handlers

import (
	"Phylogeny/database/queries"
	"Phylogeny/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

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

func GetJobHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := queries.GetJobByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("Job found with ID:", job.ID)
	return c.JSON(job)
}

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

func DeleteJobHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := queries.DeleteJob(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	log.Println("Job deleted with ID:", id)
	return c.SendStatus(fiber.StatusNoContent)
}
