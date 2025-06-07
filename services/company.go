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

type CompanyServiceInterface interface {
	CreateCompany(ctx context.Context, company *models.Company) error
	GetCompany(ctx context.Context, id uuid.UUID) (*models.Company, error)
	GetCompanyByAPIKey(ctx context.Context, apiKey string) (*models.Company, error)
	GetCompanyByCnpj(ctx context.Context, cnpj string) (*models.Company, error)
	GetAllCompanies(ctx context.Context, limit, offset int) ([]models.Company, error)
	UpdateCompany(ctx context.Context, company *models.Company) error
	DeactivateCompany(ctx context.Context, id uuid.UUID) error
	RotateAPIKey(ctx context.Context, id uuid.UUID) (string, error)
}

type CompanyService struct {
	Repo repositories.CompanyRepositoryInterface
}

var _ CompanyServiceInterface = &CompanyService{}

func (s *CompanyService) CreateCompany(ctx context.Context, company *models.Company) error {
	if company.Name == "" {
		return errors.New("company name is required")
	}
	if company.Cnpj == "" {
		return errors.New("company CNPJ is required")
	}

	company.ID = uuid.New()
	company.IsActive = true
	now := time.Now()
	company.CreatedAt = now
	company.UpdatedAt = now

	newKey, err := s.Repo.RotateAPIKey(ctx, company.ID)
	if err != nil {
		return fmt.Errorf("failed to generate initial API key: %w", err)
	}
	company.APIKey = newKey

	if err := s.Repo.CreateCompany(ctx, company); err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}
	return nil
}

func (s *CompanyService) GetCompany(ctx context.Context, id uuid.UUID) (*models.Company, error) {
	company, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get company: %w", err)
	}
	return company, nil
}

func (s *CompanyService) GetCompanyByAPIKey(ctx context.Context, apiKey string) (*models.Company, error) {
	company, err := s.Repo.FindByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get company by API key: %w", err)
	}
	return company, nil
}

func (s *CompanyService) GetCompanyByCnpj(ctx context.Context, cnpj string) (*models.Company, error) {
	company, err := s.Repo.FindByCnpj(ctx, cnpj)
	if err != nil {
		return nil, fmt.Errorf("failed to get company by CNPJ: %w", err)
	}
	return company, nil
}

func (s *CompanyService) GetAllCompanies(ctx context.Context, limit, offset int) ([]models.Company, error) {
	companies, err := s.Repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get companies: %w", err)
	}
	return companies, nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, company *models.Company) error {
	if company.Name == "" {
		return errors.New("company name is required")
	}
	if company.Cnpj == "" {
		return errors.New("company CNPJ is required")
	}
	company.UpdatedAt = time.Now()

	if err := s.Repo.UpdateCompany(ctx, company); err != nil {
		return fmt.Errorf("failed to update company: %w", err)
	}
	return nil
}

func (s *CompanyService) DeactivateCompany(ctx context.Context, id uuid.UUID) error {
	if err := s.Repo.DeactivateCompany(ctx, id); err != nil {
		return fmt.Errorf("failed to deactivate company: %w", err)
	}
	return nil
}

func (s *CompanyService) RotateAPIKey(ctx context.Context, id uuid.UUID) (string, error) {
	newKey, err := s.Repo.RotateAPIKey(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to rotate API key: %w", err)
	}
	return newKey, nil
}
