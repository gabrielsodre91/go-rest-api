package main

import (
	"log"

	"github.com/gabrielsodre91/api-gin/database"
	"github.com/gabrielsodre91/api-gin/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }
	
	database.Connect()
	
	server := server.NewServer()
	server.Run()
}