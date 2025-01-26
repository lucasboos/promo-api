package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"promo-api/config"
	"promo-api/controllers"
	"promo-api/middlewares"
	"promo-api/repositories"
	"promo-api/routes"
	"promo-api/services"
)

func main() {
	db := config.GetDB()
	defer config.CloseDB()

	repo := &repositories.PromotionRepository{DB: db}
	service := &services.PromotionService{Repo: repo}
	controller := &controllers.PromotionController{Service: service}

	r := mux.NewRouter()
	r.Use(middlewares.ValidateContentType)
	routes.ConfigureRoutes(r, controller)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
