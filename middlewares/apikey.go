package middlewares

import (
	"context"
	"net/http"
	"strings"

	"promo-api/repositories"
)

type contextKey string

const CompanyContextKey = contextKey("company")

func ValidateAPIKey(repo repositories.CompanyRepositoryInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			if strings.TrimSpace(apiKey) == "" {
				http.Error(w, "API key required", http.StatusUnauthorized)
				return
			}

			company, err := repo.FindByAPIKey(r.Context(), apiKey)
			if err != nil || company == nil || !company.IsActive {
				http.Error(w, "Invalid or inactive API key", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), CompanyContextKey, company)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
