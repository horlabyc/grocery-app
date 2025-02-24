package main

import (
	"fmt"
	"log"

	"github.com/horlabyc/grocery-app/internal/storage/postgres"
	"github.com/horlabyc/grocery-app/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// initialize db config
	dbConfig := postgres.Config{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnvAsInt("DB_PORT", 5432),
		User:     utils.GetEnv("DB_USER", "admin"),
		Password: utils.GetEnv("DB_PASSWORD", "password"),
		DBName:   utils.GetEnv("DB_NAME", "grocery_app"),
		SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
	}

	db, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()
	fmt.Println("Connected to database")
}
