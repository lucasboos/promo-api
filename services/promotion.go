package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"promo-api/models"
	"promo-api/repositories"
)

type PromotionServiceInterface interface {
	CreatePromotion(ctx context.Context, promotion *models.Promotion) error
	GetPromotion(ctx context.Context, id uuid.UUID) (*models.Promotion, error)
	GetAllPromotions(ctx context.Context, limit, offset int) ([]models.Promotion, error)
	GetPromotionsByCoupon(ctx context.Context, coupon string) ([]models.Promotion, error)
	UpdatePromotion(ctx context.Context, promotion *models.Promotion) error
	DeletePromotion(ctx context.Context, id uuid.UUID) error
}

type PromotionService struct {
	Repo repositories.PromotionRepositoryInterface
}

var _ PromotionServiceInterface = &PromotionService{}

func (s *PromotionService) CreatePromotion(ctx context.Context, promotion *models.Promotion) error {
	if promotion.StartDate.After(promotion.EndDate) {
		return errors.New("start_date cannot be after end_date")
	}
	if promotion.DiscountValue <= 0 {
		return errors.New("discount_value must be greater than zero")
	}

	promotion.ID = uuid.New()
	now := time.Now()
	promotion.CreatedAt = now
	promotion.UpdatedAt = now

	if err := s.Repo.CreatePromotion(ctx, promotion); err != nil {
		return fmt.Errorf("failed to create promotion: %w", err)
	}
	return nil
}

func (s *PromotionService) GetAllPromotions(ctx context.Context, limit, offset int) ([]models.Promotion, error) {
	promotions, err := s.Repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get promotions: %w", err)
	}
	return promotions, nil
}

func (s *PromotionService) GetPromotion(ctx context.Context, id uuid.UUID) (*models.Promotion, error) {
	promotion, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get promotion: %w", err)
	}
	return promotion, nil
}

func (s *PromotionService) GetPromotionsByCoupon(ctx context.Context, coupon string) ([]models.Promotion, error) {
	promotions, err := s.Repo.FindByCoupon(ctx, coupon)
	if err != nil {
		return nil, fmt.Errorf("failed to get promotion by coupon: %w", err)
	}
	return promotions, nil
}

func (s *PromotionService) UpdatePromotion(ctx context.Context, promotion *models.Promotion) error {
	if promotion.StartDate.After(promotion.EndDate) {
		return errors.New("start_date cannot be after end_date")
	}
	if promotion.DiscountValue <= 0 {
		return errors.New("discount_value must be greater than zero")
	}

	promotion.UpdatedAt = time.Now()

	if err := s.Repo.UpdatePromotion(ctx, promotion); err != nil {
		return fmt.Errorf("failed to update promotion: %w", err)
	}
	return nil
}

func (s *PromotionService) DeletePromotion(ctx context.Context, id uuid.UUID) error {
	if err := s.Repo.DeletePromotion(ctx, id); err != nil {
		return fmt.Errorf("failed to delete promotion: %w", err)
	}
	return nil
}
