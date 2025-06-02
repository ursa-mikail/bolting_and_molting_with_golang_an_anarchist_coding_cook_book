package config

import (
	"fmt"
	"log"

	env "github.com/caarlos0/env/v11"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string `env:"APP_PORT"`
	DBHost  string `env:"DB_HOST"`
	DBUser  string `env:"DB_USER"`
	DBPass  string `env:"DB_PASS"`
}

func LoadConfig(envFile string) (Config, error) {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: Could not load .env file %s: %v", envFile, err)
		return Config{}, fmt.Errorf("Failed to load .env file: %w", err)
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Printf("Warning: Could not load .env file %s: %v", envFile, err)
		return Config{}, fmt.Errorf("Failed to load .env file: %w", err)
	}

	log.Println("Config loaded success")

	return cfg, nil
}
