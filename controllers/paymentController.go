package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/database"
	"project/models"
	"time"

	"github.com/gorilla/mux"
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

// GetAllPayments retrieves all payment details from the database
func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, date, selected_items, total, tax, payable FROM payments")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching payments: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		// Update the scanning of the 'date' field
		var date []byte
		err := rows.Scan(&payment.ID, &date, &payment.SelectedItems, &payment.Total, &payment.Tax, &payment.Payable)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}

		// Convert the date []byte to time.Time
		if len(date) > 0 {
			payment.Date, err = time.Parse("2006-01-02 15:04:05", string(date))
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing date: %v", err), http.StatusInternalServerError)
				return
			}
		}
		payments = append(payments, payment)
	}

	// Send the list of payments as JSON
	json.NewEncoder(w).Encode(payments)
}

// GetPaymentByID retrieves a single payment's details by its ID
func GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	// Extract the payment ID from the URL
	params := mux.Vars(r)
	id := params["id"]

	// Query the database for payment details by ID
	var payment models.Payment
	var date []byte
	query := "SELECT id, date, selected_items, total, tax, payable FROM payments WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&payment.ID, &date, &payment.SelectedItems, &payment.Total, &payment.Tax, &payment.Payable)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching payment: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the date []byte to time.Time
	if len(date) > 0 {
		payment.Date, err = time.Parse("2006-01-02 15:04:05", string(date))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing date: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Send the payment details as JSON
	json.NewEncoder(w).Encode(payment)
}
