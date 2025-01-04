package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/database"
	"project/models"
	"time"
)

// HandleCheckout stores the checkout details in the database
func HandleCheckout(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	// Decode the incoming JSON body
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Set the current time as the date for the transaction
	payment.Date = time.Now()

	// Save the payment to the database
	err = storePayment(payment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving payment: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Payment details saved successfully!",
	})
}

// storePayment saves the payment details in the database
func storePayment(payment models.Payment) error {
	query := "INSERT INTO payments (date, selected_items, total, tax, payable) VALUES (?, ?, ?, ?, ?)"

	_, err := database.DB.Exec(query, payment.Date, payment.SelectedItems, payment.Total, payment.Tax, payment.Payable)
	if err != nil {
		return fmt.Errorf("could not insert payment into database: %v", err)
	}

	return nil
}
