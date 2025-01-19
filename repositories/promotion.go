package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"promo-api/models"
)

type PromotionRepository struct {
	DB *sqlx.DB
}

func (r *PromotionRepository) CreatePromotion(promotion *models.Promotion) error {
	query := `
		INSERT INTO promotions (id, title, description, discount_type, discount_value, start_date, end_date, minimum_purchase_amount, max_usage, current_usage, coupon_code, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err := r.DB.Exec(query,
		promotion.ID, promotion.Title, promotion.Description, promotion.DiscountType, promotion.DiscountValue,
		promotion.StartDate, promotion.EndDate, promotion.MinimumPurchaseAmount, promotion.MaxUsage,
		promotion.CurrentUsage, promotion.CouponCode, promotion.IsActive, promotion.CreatedAt, promotion.UpdatedAt,
	)
	return err
}

func (r *PromotionRepository) GetPromotionByID(id uuid.UUID) (*models.Promotion, error) {
	var promotion models.Promotion
	query := "SELECT * FROM promotions WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(
		&promotion.ID, &promotion.Title, &promotion.Description, &promotion.DiscountType, &promotion.DiscountValue,
		&promotion.StartDate, &promotion.EndDate, &promotion.MinimumPurchaseAmount, &promotion.MaxUsage,
		&promotion.CurrentUsage, &promotion.CouponCode, &promotion.IsActive, &promotion.CreatedAt, &promotion.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &promotion, nil
}

func (r *PromotionRepository) GetAllPromotions() ([]models.Promotion, error) {
	var promotions []models.Promotion
	query := "SELECT * FROM promotions"
	err := r.DB.Select(&promotions, query)
	if err != nil {
		return nil, err
	}
	return promotions, nil
}
