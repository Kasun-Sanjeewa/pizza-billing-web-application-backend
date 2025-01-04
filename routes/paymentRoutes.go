package routes

import (
	"github.com/gorilla/mux"
	"project/controllers"
)

func RegisterPaymentRoutes(router *mux.Router) {
	router.HandleFunc("/checkout", controllers.HandleCheckout).Methods("POST")
}
