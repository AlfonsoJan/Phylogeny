package main

import (
	"Phylogeny/config"
	"Phylogeny/database"
	"Phylogeny/routes"
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	database.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
