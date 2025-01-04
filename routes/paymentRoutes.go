package routes

import (
	"project/controllers"

	"github.com/gorilla/mux"
)

func RegisterPaymentRoutes(router *mux.Router) {
	// Route for checking out and storing the payment
	router.HandleFunc("/checkout", controllers.HandleCheckout).Methods("POST")

	// Route to get all payment details
	router.HandleFunc("/payments", controllers.GetAllPayments).Methods("GET")

}
