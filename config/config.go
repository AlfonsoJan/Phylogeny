package config

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func GetDatabaseURL() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		panic("One or more database environment variables are missing. Please check .env file.")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

func GetEnv() *string {
	env := flag.String("env", "dev", "Set the environment (dev or prod)")
	flag.Parse()
	if *env != "dev" && *env != "prod" {
		log.Fatalf("Invalid environment value: %s. Only 'dev' or 'prod' are allowed.", *env)
	}
	os.Setenv("ENV", *env)
	log.Println("Environment is", *env)
	return env
}
