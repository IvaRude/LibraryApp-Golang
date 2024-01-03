package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func NewConfig() *Config {
	return &Config{
		Host:     getEnv("host", "localhost"),
		Port:     getEnv("port", "5432"),
		User:     getEnv("user", "postgres"),
		Password: getEnv("password", "postgres"),
		DbName:   getEnv("dbname", "postgres"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
