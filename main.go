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

	companyRepo := &repositories.CompanyRepository{DB: db}
	companyService := &services.CompanyService{Repo: companyRepo}
	companyController := &controllers.CompanyController{Service: companyService}

	promoRepo := &repositories.PromotionRepository{DB: db}
	promoService := &services.PromotionService{Repo: promoRepo}
	promoController := &controllers.PromotionController{Service: promoService}

	r := mux.NewRouter()
	r.Use(middlewares.ValidateContentType)

	authorized := r.PathPrefix("/").Subrouter()
	authorized.Use(middlewares.ValidateAPIKey(companyRepo))

	routes.ConfigurePromotionRoutes(authorized, promoController)
	routes.ConfigureCompanyRoutes(authorized, companyController)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
