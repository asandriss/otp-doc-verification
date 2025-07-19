package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	// ..
}

func LoadConfig(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env found, falling back to env variables")
		return err
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8085" // using 85 to clearly see if config load failed
	}

	cfg.Port = port

	return nil
}
