package main

import (
	"fmt"
	"log"
	"os"
	"people-api/cmd/server"
	db "people-api/internal/app/database"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	//Dotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Init logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Error in logger, %s", err)
	}
	defer logger.Sync()

	//Init db
	_, errdb := db.GetInstance()

	if errdb != nil {
		fmt.Println("Error creatin the db connection")
		os.Exit(1)
	} else {
		fmt.Println("DB connected...")
	}

	// initialice the server
	errServer := server.Server(logger)
	if errServer != nil {
		os.Exit(2)
	}

}
