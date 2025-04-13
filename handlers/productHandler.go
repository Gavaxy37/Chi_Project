package handlers

import (
	"Chi_Project/db"
	"Chi_Project/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	db.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewDecoder(r.Body).Decode(&product)
	db.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	db.DB.Delete(&models.Product{}, id)
	w.WriteHeader(http.StatusOK)
}
