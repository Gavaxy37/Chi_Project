package main

import (
	"Chi_Project/config"
	"Chi_Project/db"
	_ "Chi_Project/docs"
	"Chi_Project/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title Chi API
// @version 1.0
// @description A simple product API with Chi + PostgreSQL
// @host localhost:8080
// @BasePath /

func main() {
	config.LoadConfig()

	// log level
	level, err := logrus.ParseLevel(config.AppConfig.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}

	logrus.Infof("Starting with log level: %s", config.AppConfig.LogLevel)

	db.Connect()

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/products", handlers.GetProducts)
	r.Post("/products", handlers.CreateProduct)
	r.Get("/products/{id}", handlers.GetProduct)
	r.Put("/products/{id}", handlers.UpdateProduct)
	r.Delete("/products/{id}", handlers.DeleteProduct)

	logrus.Info("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
