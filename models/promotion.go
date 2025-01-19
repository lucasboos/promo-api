package models

import (
	"time"

	"github.com/google/uuid"
)

type Promotion struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	Title               string    `json:"title" db:"title"`
	Description         string    `json:"description,omitempty" db:"description"`
	DiscountType        string    `json:"discount_type" db:"discount_type"`
	DiscountValue       float64   `json:"discount_value" db:"discount_value"`
	StartDate           time.Time `json:"start_date" db:"start_date"`
	EndDate             time.Time `json:"end_date" db:"end_date"`
	MinimumPurchaseAmount *float64  `json:"minimum_purchase_amount,omitempty" db:"minimum_purchase_amount"`
	MaxUsage            *int      `json:"max_usage,omitempty" db:"max_usage"`
	CurrentUsage        int       `json:"current_usage" db:"current_usage"`
	CouponCode          *string   `json:"coupon_code,omitempty" db:"coupon_code"`
	IsActive            bool      `json:"is_active" db:"is_active"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
