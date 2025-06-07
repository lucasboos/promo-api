package repositories

import (
	"context"
	"fmt"
	"time"

	"promo-api/models"

	"promo-api/utils"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompanyRepositoryInterface interface {
	CreateCompany(ctx context.Context, company *models.Company) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Company, error)
	FindByAPIKey(ctx context.Context, apiKey string) (*models.Company, error)
	FindByCnpj(ctx context.Context, cnpj string) (*models.Company, error)
	FindAll(ctx context.Context, limit, offset int) ([]models.Company, error)
	UpdateCompany(ctx context.Context, company *models.Company) error
	DeactivateCompany(ctx context.Context, id uuid.UUID) error
	RotateAPIKey(ctx context.Context, id uuid.UUID) (string, error)
}

type CompanyRepository struct {
	DB *sqlx.DB
}

var _ CompanyRepositoryInterface = &CompanyRepository{}

func (r *CompanyRepository) CreateCompany(ctx context.Context, company *models.Company) error {
	query := `
		INSERT INTO companies (
			id, name, cnpj, api_key, is_active, created_at, updated_at, deleted_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, NULL
		)`
	_, err := r.DB.ExecContext(ctx, query,
		company.ID, company.Name, company.Cnpj, company.APIKey,
		company.IsActive, company.CreatedAt, company.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}
	return nil
}

func (r *CompanyRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Company, error) {
	var company models.Company
	query := "SELECT * FROM companies WHERE id = $1 AND deleted_at IS NULL"
	err := r.DB.GetContext(ctx, &company, query, id)
	if err != nil {
		return nil, fmt.Errorf("company not found with ID %s: %w", id, err)
	}
	return &company, nil
}

func (r *CompanyRepository) FindByAPIKey(ctx context.Context, apiKey string) (*models.Company, error) {
	var company models.Company
	query := "SELECT * FROM companies WHERE api_key = $1 AND deleted_at IS NULL"
	err := r.DB.GetContext(ctx, &company, query, apiKey)
	if err != nil {
		return nil, fmt.Errorf("company not found with API key: %w", err)
	}
	return &company, nil
}

func (r *CompanyRepository) FindByCnpj(ctx context.Context, cnpj string) (*models.Company, error) {
	var company models.Company
	query := "SELECT * FROM companies WHERE cnpj = $1 AND deleted_at IS NULL"
	err := r.DB.GetContext(ctx, &company, query, cnpj)
	if err != nil {
		return nil, fmt.Errorf("company not found with CNPJ %s: %w", cnpj, err)
	}
	return &company, nil
}

func (r *CompanyRepository) FindAll(ctx context.Context, limit, offset int) ([]models.Company, error) {
	var companies []models.Company
	query := "SELECT * FROM companies WHERE deleted_at IS NULL LIMIT $1 OFFSET $2"
	err := r.DB.SelectContext(ctx, &companies, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch companies: %w", err)
	}
	return companies, nil
}

func (r *CompanyRepository) UpdateCompany(ctx context.Context, company *models.Company) error {
	query := `
		UPDATE companies
		SET name = $1, cnpj = $2, api_key = $3, is_active = $4, updated_at = $5
		WHERE id = $6 AND deleted_at IS NULL`
	_, err := r.DB.ExecContext(ctx, query,
		company.Name, company.Cnpj, company.APIKey, company.IsActive, company.UpdatedAt, company.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update company %s: %w", company.Name, err)
	}
	return nil
}

func (r *CompanyRepository) DeactivateCompany(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE companies
		SET is_active = false, deleted_at = $1, updated_at = $1
		WHERE id = $2 AND deleted_at IS NULL`
	now := time.Now()
	_, err := r.DB.ExecContext(ctx, query, now, id)
	if err != nil {
		return fmt.Errorf("failed to deactivate company with ID %s: %w", id, err)
	}
	return nil
}

func (r *CompanyRepository) RotateAPIKey(ctx context.Context, id uuid.UUID) (string, error) {
	newKey, err := utils.GenerateAPIKey()
	if err != nil {
		return "", fmt.Errorf("failed to generate new API key: %w", err)
	}

	query := `
		UPDATE companies
		SET api_key = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL`
	_, err = r.DB.ExecContext(ctx, query, newKey, time.Now(), id)
	if err != nil {
		return "", fmt.Errorf("failed to rotate API key for company %s: %w", id, err)
	}

	return newKey, nil
}
