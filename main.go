package main

import (
	"Chi_Project/db"
	"Chi_Project/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	r := chi.NewRouter()

	r.Get("/products", handlers.GetProducts)
	r.Post("/products", handlers.CreateProduct)
	r.Get("/products/{id}", handlers.GetProduct)
	r.Put("/products/{id}", handlers.UpdateProduct)
	r.Delete("/products/{id}", handlers.DeleteProduct)

	log.Println("Server on :8080")
	http.ListenAndServe(":8080", r)
}
