package main

import (
	"imageupload/app"
	_ "imageupload/docs"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	//load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnv()
	app.Routes()
}
