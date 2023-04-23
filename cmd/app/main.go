package main

import (
	"github.com/joho/godotenv"
	"log"
	"smth/config"
	"smth/internal/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	conf := config.New()

	if err := app.Start(conf); err != nil {
		log.Fatal(err)
	}
}
