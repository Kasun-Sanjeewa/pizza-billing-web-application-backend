package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"project/database"
	"project/models"

	"github.com/gorilla/mux"
)

// Create a product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	query := "INSERT INTO products (name, barcode, price, img) VALUES (?, ?, ?, ?)"
	res, err := database.DB.Exec(query, product.Name, product.Barcode, product.Price, product.Img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	product.ID = int(id)
	json.NewEncoder(w).Encode(product)
}

// Get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, barcode, price, img FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(&product.ID, &product.Name, &product.Barcode, &product.Price, &product.Img)
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

// Get a product by ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var product models.Product
	query := "SELECT id, name, barcode, price, img FROM products WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Barcode, &product.Price, &product.Img)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// Update a product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	query := "UPDATE products SET name = ?, barcode = ?, price = ?, img = ? WHERE id = ?"
	_, err := database.DB.Exec(query, product.Name, product.Barcode, product.Price, product.Img, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ID = id
	json.NewEncoder(w).Encode(product)
}

// Delete a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := "DELETE FROM products WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}
