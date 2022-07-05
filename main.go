package main

import (
	"blog-api/config"
	"blog-api/src"
	"log"
	"sync"

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

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
