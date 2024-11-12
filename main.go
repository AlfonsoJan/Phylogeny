package main

import (
	"Phylogeny/config"
	"Phylogeny/setup"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var err error
	err = config.LoadEnv()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
	app := fiber.New()

	err = setup.Setup(app)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
}
