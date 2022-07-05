package main

import (
	"blog-api/config"
	"blog-api/src"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := config.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}

	server := src.InitServer(db)

	server.Run()
}
