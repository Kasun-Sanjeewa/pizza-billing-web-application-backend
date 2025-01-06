package routes

import (
	"project/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
}
