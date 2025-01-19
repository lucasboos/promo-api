package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"promo-api/models"
	"promo-api/services"
)

type PromotionController struct {
	Service *services.PromotionService
}

func (c *PromotionController) CreatePromotion(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	if err := json.NewDecoder(r.Body).Decode(&promotion); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	promotion.ID = uuid.New()
	promotion.CreatedAt = time.Now()
	promotion.UpdatedAt = time.Now()

	if err := c.Service.CreatePromotion(&promotion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promotion)
}

func (c *PromotionController) GetPromotion(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	promotion, err := c.Service.GetPromotion(id)
	if err != nil {
		http.Error(w, "Promotion not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promotion)
}

func (c *PromotionController) GetAllPromotions(w http.ResponseWriter, r *http.Request) {
	promotions, err := c.Service.GetAllPromotions()
	if err != nil {
		http.Error(w, "Failed to get promotions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(promotions); err != nil {
		http.Error(w, "Failed to encode promotions", http.StatusInternalServerError)
	}
}
