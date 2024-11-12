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
	setupRoutes(app)
	setupMiddleWares(app)
	// Always the last function to be called
	startServerWithGracefulShutdown(app)
	return nil
}

func setupMiddleWares(app *fiber.App) {
	app.Use(middleware.WebApiLogger)
}

func startServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
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
	api := app.Group("/api/v1")
	api.Post("/job", handlers.CreateJobHandler)

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
