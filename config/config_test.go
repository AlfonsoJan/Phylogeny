package config

import (
	"os"
	"testing"
)

func TestGetDatabaseURLMissingVariables(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	GetDatabaseURL()
}

func TestGetDatabaseURL(t *testing.T) {
	err := os.Setenv("DB_HOST", "localhost")
	if err != nil {
		t.Fatalf("Error setting DB_HOST environment variable: %v", err)
	}
	err = os.Setenv("DB_USER", "user")
	if err != nil {
		t.Fatalf("Error setting DB_USER environment variable: %v", err)
	}
	err = os.Setenv("DB_PASSWORD", "password")
	if err != nil {
		t.Fatalf("Error setting DB_PASSWORD environment variable: %v", err)
	}
	err = os.Setenv("DB_NAME", "testdb")
	if err != nil {
		t.Fatalf("Error setting DB_NAME environment variable: %v", err)
	}
	err = os.Setenv("DB_PORT", "5432")
	if err != nil {
		t.Fatalf("Error setting DB_PORT environment variable: %v", err)
	}

	url := GetDatabaseURL()

	expectedURL := "host=localhost user=user password=password dbname=testdb port=5432 sslmode=disable"

	if url != expectedURL {
		t.Errorf("Expected %s but got %s", expectedURL, url)
	}
}
