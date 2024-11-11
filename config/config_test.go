package config

import (
	"flag"
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

func TestGetEnv(t *testing.T) {
	tests := []struct {
		args          []string
		expectedEnv   string
		expectedError bool
	}{
		{
			args:          []string{"-env", "dev"},
			expectedEnv:   "dev",
			expectedError: false,
		},
		{
			args:          []string{"-env", "prod"},
			expectedEnv:   "prod",
			expectedError: false,
		},
		{
			args:          []string{"-env", "staging"},
			expectedEnv:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		os.Args = append([]string{"cmd"}, tt.args...)

		os.Unsetenv("ENV")

		env, err := GetEnv()

		if (err != nil) != tt.expectedError {
			t.Errorf("GetEnv() error = %v, wantError %v", err, tt.expectedError)
			continue
		}

		if err == nil && *env != tt.expectedEnv {
			t.Errorf("GetEnv() = %v, want %v", *env, tt.expectedEnv)
		}

		if tt.expectedError && os.Getenv("ENV") != "" {
			t.Errorf("os.Getenv(ENV) = %v, want empty", os.Getenv("ENV"))
		} else if !tt.expectedError && os.Getenv("ENV") != tt.expectedEnv {
			t.Errorf("os.Getenv(ENV) = %v, want %v", os.Getenv("ENV"), tt.expectedEnv)
		}
	}
}
