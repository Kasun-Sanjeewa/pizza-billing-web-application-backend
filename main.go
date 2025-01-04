package main

import (
	"log"
	"net/http"
	"project/database"
	"project/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
