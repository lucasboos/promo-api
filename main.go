package main

import (
	"net/http"
	"log"

	_ "github.com/lib/pq"

	"promo-api/controllers"
	"promo-api/repositories"
	"promo-api/services"
	"promo-api/config"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	repo := &repositories.PromotionRepository{DB: config.DB}
	service := &services.PromotionService{Repo: repo}
	controller := &controllers.PromotionController{Service: service}

	http.HandleFunc("/promotions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.GetAllPromotions(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/promotion", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.CreatePromotion(w, r)
		} else if r.Method == http.MethodGet {
			controller.GetPromotion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
