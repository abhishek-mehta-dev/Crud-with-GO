package main

import (
	"log"
	"net/http"
	"os"

	"crud-with-go/config"
	"crud-with-go/router"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	// connect DB
	config.ConnectDB()

	// start router
	r := router.SetupRouter()
	port := os.Getenv("APP_PORT")
	log.Println(" Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
