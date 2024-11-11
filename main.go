package main

import (
	"Phylogeny/config"
	"Phylogeny/database"
	"Phylogeny/routes"
	"log"

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

	database.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
