package main

import (
	"Phylogeny/config"
	"Phylogeny/database"
	"Phylogeny/handlers"
	"Phylogeny/middleware"
	"Phylogeny/routes"
	"Phylogeny/tasks"
	"Phylogeny/utils"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {
	env, err := config.GetEnv()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}
	if err := godotenv.Load(".env." + *env); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))
	app.Use(middleware.WebApiLogger)

	database.Connect()

	handlers.JobQueue = tasks.NewJobQueue(10)
	handlers.JobQueue.StartWorkers(3)

	tempDir := filepath.Join(os.TempDir(), "phylogeny")
	cleanupDuration := 24 * time.Hour
	utils.CleanupOldFiles(tempDir, cleanupDuration)
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		utils.CleanupOldFiles(tempDir, cleanupDuration)
	})
	c.Start()
	defer c.Stop()

	routes.SetupRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
