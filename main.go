package main

import (
	"log"
	"net/http"
	"xyz_multifinance/database"
	"xyz_multifinance/routes"
)

func main() {
	// Initialize DB
	database.InitDB()

	// Initialize routes
	router := routes.InitializeRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":7000", router))
}
