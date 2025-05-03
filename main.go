package main

import (
	"log"
	"os"

	"github.com/AnthonyThomasson/go-scratchpad/schema"
	"github.com/AnthonyThomasson/go-scratchpad/server"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	err = dbConn.AutoMigrate(schema.Clue{}, schema.Location{}, schema.Suspect{}, schema.Murder{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	server.NewServer(llm, dbConn).Start()
}
