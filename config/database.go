package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct{}

func (dbcfg *DatabaseConfig) ConfigLoad() map[string]string {
	dir, _ := os.Getwd()
	projectPath := filepath.Join(filepath.Dir(dir), filepath.Base(dir))
	filename := filepath.Join(projectPath, "/.env")

	var err = godotenv.Load(filename)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"username": os.Getenv("DB_USERNAME"),
		"password": os.Getenv("DB_PASSWORD"),
		"dbname":   os.Getenv("DB_DATABASE"),
	}
}
