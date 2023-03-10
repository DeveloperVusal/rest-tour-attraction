package core

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Helpers struct{}

func (h *Helpers) LoadEnv() {
	dir, _ := os.Getwd()
	projectPath := filepath.Join(filepath.Dir(dir), filepath.Base(dir))
	filename := filepath.Join(projectPath, "/.env")

	var err = godotenv.Load(filename)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (h *Helpers) Env(env string) string {
	return os.Getenv(env)
}
