package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	SecretKey   string
}

func Load() *Config {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := "db" // имя сервиса в docker-compose
	port := "5432"

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, db)

	return &Config{
		Port:        getEnv("PORT", "3000"),
		DatabaseURL: databaseURL,
		SecretKey:   getEnv("SECRET_TOKEN", "supersecret"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
