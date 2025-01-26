package repositories

import (
	"context"
	"fmt"

	"promo-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PromotionRepositoryInterface interface {
	CreatePromotion(ctx context.Context, promotion *models.Promotion) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Promotion, error)
	FindAll(ctx context.Context, limit, offset int) ([]models.Promotion, error)
	FindByCoupon(ctx context.Context, coupon string) ([]models.Promotion, error)
	UpdatePromotion(ctx context.Context, promotion *models.Promotion) error
	DeletePromotion(ctx context.Context, id uuid.UUID) error
}

type PromotionRepository struct {
	DB *sqlx.DB
}

var _ PromotionRepositoryInterface = &PromotionRepository{}

func (r *PromotionRepository) CreatePromotion(ctx context.Context, promotion *models.Promotion) error {
	query := `
		INSERT INTO promotions (
			id, title, description, discount_type, discount_value, start_date, end_date,
			minimum_purchase_amount, max_usage, current_usage, coupon_code, is_active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
		)`
	_, err := r.DB.ExecContext(ctx, query,
		promotion.ID, promotion.Title, promotion.Description, promotion.DiscountType, promotion.DiscountValue,
		promotion.StartDate, promotion.EndDate, promotion.MinimumPurchaseAmount, promotion.MaxUsage,
		promotion.CurrentUsage, promotion.CouponCode, promotion.IsActive, promotion.CreatedAt, promotion.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create promotion: %w", err)
	}
	return nil
}

func (r *PromotionRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Promotion, error) {
	var promotion models.Promotion
	query := "SELECT * FROM promotions WHERE id = $1"
	err := r.DB.GetContext(ctx, &promotion, query, id)
	if err != nil {
		return nil, fmt.Errorf("promotion not found with ID %s: %w", id, err)
	}
	return &promotion, nil
}

func (r *PromotionRepository) FindAll(ctx context.Context, limit, offset int) ([]models.Promotion, error) {
	var promotions []models.Promotion
	query := "SELECT * FROM promotions LIMIT $1 OFFSET $2"
	err := r.DB.SelectContext(ctx, &promotions, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch promotions: %w", err)
	}
	return promotions, nil
}

func (r *PromotionRepository) FindByCoupon(ctx context.Context, coupon string) ([]models.Promotion, error) {
	var promotions []models.Promotion
	query := `
		SELECT * FROM promotions 
		WHERE coupon_code ILIKE $1`
	err := r.DB.SelectContext(ctx, &promotions, query, "%"+coupon+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch promotions by filter: %w", err)
	}
	return promotions, nil
}

func (r *PromotionRepository) UpdatePromotion(ctx context.Context, promotion *models.Promotion) error {
	query := `
		UPDATE promotions
		SET title = $1, description = $2, discount_type = $3, discount_value = $4,
			start_date = $5, end_date = $6, minimum_purchase_amount = $7, max_usage = $8,
			current_usage = $9, coupon_code = $10, is_active = $11, updated_at = $12
		WHERE id = $13`
	_, err := r.DB.ExecContext(ctx, query,
		promotion.Title, promotion.Description, promotion.DiscountType, promotion.DiscountValue,
		promotion.StartDate, promotion.EndDate, promotion.MinimumPurchaseAmount, promotion.MaxUsage,
		promotion.CurrentUsage, promotion.CouponCode, promotion.IsActive, promotion.UpdatedAt, promotion.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update promotion %s: %w", promotion.Title, err)
	}
	return nil
}

func (r *PromotionRepository) DeletePromotion(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM promotions WHERE id = $1"
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete promotion with ID %s: %w", id, err)
	}
	return nil
}
