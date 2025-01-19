package services

import (
	"errors"

	"github.com/google/uuid"

	"promo-api/models"
	"promo-api/repositories"
)

type PromotionService struct {
	Repo *repositories.PromotionRepository
}

func (s *PromotionService) CreatePromotion(promotion *models.Promotion) error {
	if promotion.StartDate.After(promotion.EndDate) {
		return errors.New("start_date cannot be after end_date")
	}
	if promotion.DiscountValue <= 0 {
		return errors.New("discount_value must be greater than zero")
	}
	return s.Repo.CreatePromotion(promotion)
}

func (s *PromotionService) GetPromotion(id uuid.UUID) (*models.Promotion, error) {
	return s.Repo.GetPromotionByID(id)
}

func (s *PromotionService) GetAllPromotions() ([]models.Promotion, error) {
	return s.Repo.GetAllPromotions()
}
