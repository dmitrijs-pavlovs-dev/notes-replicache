package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    PostgresUser     string
    PostgresPassword string
    PostgresDB       string
    PostgresHost     string
    PostgresPort     string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    config := &Config{
        PostgresUser:     os.Getenv("POSTGRES_USER"),
        PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
        PostgresDB:       os.Getenv("POSTGRES_DB"),
        PostgresHost:     os.Getenv("POSTGRES_HOST"),
        PostgresPort:     os.Getenv("POSTGRES_PORT"),
    }
    return config
}
