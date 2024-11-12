package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
	PORT  string
	DEV   bool
}

var EnvConfig *Config

func getDatabaseURL() (string, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		return "", errors.New("one or more database environment variables are missing. Please check .env file")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port), nil
}

func getPort() (string, error) {
	port := os.Getenv("SERVER_URL")
	if port == "" {
		return "", fmt.Errorf("SERVER_URL is not set in .env file")
	}
	return port, nil
}

func LoadEnv() error {
	env := flag.String("env", "dev", "Set the environment (dev or prod)")
	flag.Parse()
	if *env != "dev" && *env != "prod" {
		return fmt.Errorf("invalid environment: %s. Please use 'dev' or 'prod'", *env)
	}
	log.Println("Enviroment is set to", *env)
	if err := godotenv.Load(".env." + *env); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	dbURL, err := getDatabaseURL()
	if err != nil {
		return err
	}
	serverPort, err := getPort()
	if err != nil {
		return err
	}
	EnvConfig = &Config{
		DBURL: dbURL,
		PORT:  serverPort,
		DEV:   *env == "dev",
	}
	return nil
}
