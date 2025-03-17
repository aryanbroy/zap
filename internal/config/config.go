package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CLIENT_ID     string
	CLIENT_SECRET string
	PORT          string
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file", err)
		return nil
	}
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	port := os.Getenv("PORT")

	response := Config{
		CLIENT_ID:     clientId,
		CLIENT_SECRET: clientSecret,
		PORT:          port,
	}
	return &response
}
