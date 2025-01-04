package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"project/database"
	"project/routes"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new router
	r := mux.NewRouter()

	// Initialize routes
	routes.RegisterProductRoutes(r)

	// CORS configuration
	corsAllowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	// Start the server with CORS enabled
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(r)))
}
