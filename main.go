package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rennanbadaro/discount-calculator/cmd/grpc"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "local" || appEnv == "test" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("failed to load .env file: %v", err)
		}
	}

	if err := grpc.StartServer(); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
