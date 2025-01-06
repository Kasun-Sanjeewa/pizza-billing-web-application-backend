package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"project/database"
	"project/models"
)

// LoginHandler handles the login logic
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate the user credentials
	if validateUser(user.Username, user.Password) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

// validateUser checks if the username and password exist in the database
func validateUser(username, password string) bool {
	var dbPassword string

	// Query the database
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // User not found
		}
		return false // Other errors
	}

	// Compare passwords
	return dbPassword == password
}
