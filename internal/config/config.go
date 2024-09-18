package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	// Connection strings
	ServerAddr  string
	FullServerAddr  string
	PostgresURL string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	//Выполниться один раз (для себя сингл-тон)
	once.Do(func ()  {
		log.Println("read application configuration")
		instance = &Config{
			ServerAddr:  getEnv("SERVER_ADDR", "localhost:8080"),
			FullServerAddr:  getEnv("FULL_SERVER_ADDR", "http://localhost:8080"),
			PostgresURL: getEnv("POSTGRES_URL", "postgres://test_user:test_password@postgres:5432/task_swoyo_db"),
		}
	})
	return instance
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