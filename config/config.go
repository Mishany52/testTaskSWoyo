package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Connection strings
	ServerAddr  string
	FullServerAddr  string
	PostgresURL string
}

func New() *Config {
	return &Config{
		ServerAddr:  getEnv("SERVER_ADDR", "localhost:8080"),
		FullServerAddr:  getEnv("FULL_SERVER_ADDR", "http://localhost:8080"),
		PostgresURL: getEnv("POSTGRES_URL", "postgres://test_user:test_password@postgres:5432/task_swoyo_db"),
	}
}

func getEnv(key string, defaultVal string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}