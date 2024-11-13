package setup

import (
	"Phylogeny/config"
	"Phylogeny/database"
	"Phylogeny/handlers"
	"Phylogeny/middleware"
	"Phylogeny/routes"
	"Phylogeny/tasks"
	"Phylogeny/tasks/scheduled"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/robfig/cron"
)

func Setup(app *fiber.App) error {
	err := database.Connect()
	if err != nil {
		return err
	}
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))
	setupJobQueue()
	setupScheduledTasks()
	setupMiddleWares(app)
	setupRoutes(app)
	// Always the last function to be called
	startServerWithGracefulShutdown(app)
	return nil
}

func setupMiddleWares(app *fiber.App) {
	app.Use(favicon.New(favicon.Config{
		File: "./static/images/favicon.ico",
		URL:  "/favicon.ico",
	}))
	app.Use(middleware.WebApiLogger)
}

func startServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Initiating server shutdown...")

		fiberShutdownDone := make(chan struct{})

		go func() {
			if err := a.Shutdown(); err != nil {
				log.Printf("Oops... Server shutdown failed! Reason: %v", err)
			} else {
				log.Println("Server shutdown completed.")
			}
			close(fiberShutdownDone)
		}()

		select {
		case <-fiberShutdownDone:
			log.Println("Fiber server has shut down gracefully.")
		case <-time.After(10 * time.Second):
			log.Println("Fiber shutdown timed out. Continuing with job queue shutdown.")
		}

		handlers.JobQueue.Shutdown()

		close(idleConnsClosed)
	}()
	if err := a.Listen(config.EnvConfig.PORT); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server shut down gracefully.")
}

func setupRoutes(app *fiber.App) {
	// app.Static("/", "./static/index.html")
	app.Static("/", "./static")
	api := app.Group("/api/v1")
	api.Post("/job", handlers.CreateJobHandler)
	app.Get("/ws/jobstatus/:jobID", handlers.JobStatusWebSocket(handlers.JobQueue))

	routes.NotFoundRoute(app)
}

func setupJobQueue() {
	handlers.JobQueue = tasks.NewJobQueue(10)
	handlers.JobQueue.StartWorkers(5)
}

func setupScheduledTasks() {
	c := cron.New()

	tempDir := filepath.Join(os.TempDir(), "phylogeny")
	cleanupDurationOnceADay := 24 * time.Hour
	addToSchedule(c, "0 0 * * *", func() {
		scheduled.CleanupOldFiles(tempDir, cleanupDurationOnceADay)
	})
	c.Start()
	defer c.Stop()
}

func addToSchedule(c *cron.Cron, cronExpr string, task func()) {
	task()
	err := c.AddFunc(cronExpr, task)
	if err != nil {
		log.Fatalf("Error adding task to schedule: %v", err)
	}
}
