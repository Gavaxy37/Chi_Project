package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Chi_Project/db"
	"Chi_Project/models"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// GetProducts отримує список всіх продуктів
// @Summary Отримати всі продукти
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if result := db.DB.Find(&products); result.Error != nil {
		logrus.WithError(result.Error).Error("Не вдалося отримати продукти")
		respondWithError(w, http.StatusInternalServerError, "Помилка при отриманні продуктів")
		return
	}
	json.NewEncoder(w).Encode(products)
}

// CreateProduct створює новий продукт
// @Summary Створити продукт
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Дані продукту"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logrus.WithError(err).Warn("Невірний JSON при створенні продукту")
		respondWithError(w, http.StatusBadRequest, "Невірні дані JSON")
		return
	}

	if result := db.DB.Create(&product); result.Error != nil {
		logrus.WithError(result.Error).Error("Помилка при створенні продукту в базі")
		respondWithError(w, http.StatusInternalServerError, "Не вдалося створити продукт")
		return
	}

	logrus.WithFields(logrus.Fields{"name": product.Name, "price": product.Price}).Info("Продукт створено")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProduct отримує продукт по ID
// @Summary Отримати продукт
// @Tags products
// @Produce json
// @Param id path int true "ID продукту"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Невірний ID")
		return
	}
	var product models.Product
	if result := db.DB.First(&product, id); result.Error != nil {
		logrus.WithError(result.Error).Warn("Продукт не знайдено")
		respondWithError(w, http.StatusNotFound, "Продукт не знайдено")
		return
	}
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct оновлює дані продукту
// @Summary Оновити продукт
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID продукту"
// @Param product body models.Product true "Оновлені дані продукту"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Невірний ID")
		return
	}
	var product models.Product
	if result := db.DB.First(&product, id); result.Error != nil {
		logrus.WithError(result.Error).Warn("Продукт не знайдено")
		respondWithError(w, http.StatusNotFound, "Продукт не знайдено")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Невірні дані JSON")
		return
	}
	db.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct видаляє продукт
// @Summary Видалити продукт
// @Tags products
// @Param id path int true "ID продукту"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Невірний ID")
		return
	}
	result := db.DB.Delete(&models.Product{}, id)
	if result.RowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Продукт не знайдено")
		return
	}
	logrus.WithField("id", id).Info("Продукт видалено")
	json.NewEncoder(w).Encode(map[string]string{"message": "Продукт видалено"})
}
