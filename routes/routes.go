package routes

import (
	"net/http"

	"promo-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigurePromotionRoutes(r *mux.Router, controller *controllers.PromotionController) {
	r.HandleFunc("/promotions", controller.CreatePromotion).Methods(http.MethodPost)
	r.HandleFunc("/promotions", controller.GetAllPromotions).Methods(http.MethodGet)
	r.HandleFunc("/promotions/{id}", controller.GetPromotion).Methods(http.MethodGet)
	r.HandleFunc("/promotions/coupon", controller.GetPromotionsByCoupon).Methods(http.MethodGet)
	r.HandleFunc("/promotions/{id}", controller.UpdatePromotion).Methods(http.MethodPut)
	r.HandleFunc("/promotions/{id}", controller.DeletePromotion).Methods(http.MethodDelete)
}

func ConfigureCompanyRoutes(r *mux.Router, controller *controllers.CompanyController) {
	r.HandleFunc("/companies", controller.CreateCompany).Methods(http.MethodPost)
	r.HandleFunc("/companies", controller.GetAllCompanies).Methods(http.MethodGet)
	r.HandleFunc("/companies/{id}", controller.GetCompany).Methods(http.MethodGet)
	r.HandleFunc("/companies/{id}", controller.UpdateCompany).Methods(http.MethodPut)
	r.HandleFunc("/companies/{id}", controller.DeactivateCompany).Methods(http.MethodDelete)
	r.HandleFunc("/companies/{id}/rotate-api-key", controller.RotateAPIKey).Methods(http.MethodPost)
}
