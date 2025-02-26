package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	log.Println("Loading .env file")
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getString("DATABASE_URL", ""),
	}
}

type LogConfig struct {
	Level  int
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}
