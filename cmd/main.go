package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cake-store/internal"
	"github.com/cake-store/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &internal.App{}
	app.Initialize()

	db := database.NewMySQLDB()
	defer db.Close()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Fatal(app.Router.Start(fmt.Sprintf(":%s", port)))
}
