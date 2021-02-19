package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rennanbadaro/discount-service/cmd/grpc"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("failed to load .env file: %v\n", err)
		}
	}

	if err := grpc.StartServer(); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
