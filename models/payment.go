package models

import "time"

// Payment represents the payment details.
type Payment struct {
	ID            int       `json:"id"`
	Date          time.Time `json:"date"`
	SelectedItems string    `json:"selected_items"`
	Total         float64   `json:"total"`
	Tax           float64   `json:"tax"`
	Payable       float64   `json:"payable"`
}
