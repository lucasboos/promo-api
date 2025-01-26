package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"promo-api/models"
	"promo-api/services"
)

type PromotionController struct {
	Service services.PromotionServiceInterface
}

func (c *PromotionController) CreatePromotion(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion

	if err := json.NewDecoder(r.Body).Decode(&promotion); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreatePromotion(r.Context(), &promotion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promotion)
}

func (c *PromotionController) GetPromotion(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	promotion, err := c.Service.GetPromotion(r.Context(), id)
	if err != nil {
		http.Error(w, "Promotion not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

func (c *PromotionController) GetAllPromotions(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	promotions, err := c.Service.GetAllPromotions(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, "Failed to get promotions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotions)
}

func (c *PromotionController) GetPromotionsByCoupon(w http.ResponseWriter, r *http.Request) {
	coupon := r.URL.Query().Get("coupon")

	if coupon == "" {
		http.Error(w, "Coupon is required", http.StatusBadRequest)
		return
	}

	promotions, err := c.Service.GetPromotionsByCoupon(r.Context(), coupon)
	if err != nil {
		http.Error(w, "Failed to get promotions by coupon", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotions)
}

func (c *PromotionController) UpdatePromotion(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion

	if err := json.NewDecoder(r.Body).Decode(&promotion); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.Service.UpdatePromotion(r.Context(), &promotion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

func (c *PromotionController) DeletePromotion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := c.Service.DeletePromotion(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
