package main

import (
	"log"
	"net/http"

	"self-improvement/zoom-clone-backend/config"
	"self-improvement/zoom-clone-backend/routes"
)

func main() {
	// Load configuration
	config.LoadEnv()

	config.ConnectDatabase()

	config.RunMigrations()


	// Set up routes
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe("", router))
}
