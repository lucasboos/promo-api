package routes

import (
	"net/http"

	"promo-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router, controller *controllers.PromotionController) {
	r.HandleFunc("/promotion", controller.CreatePromotion).Methods(http.MethodPost)
	r.HandleFunc("/promotions", controller.GetAllPromotions).Methods(http.MethodGet)
	r.HandleFunc("/promotion", controller.GetPromotion).Methods(http.MethodGet)
	r.HandleFunc("/promotions/coupon", controller.GetPromotionsByCoupon).Methods(http.MethodGet)
	r.HandleFunc("/promotion/{id}", controller.UpdatePromotion).Methods(http.MethodPut)
	r.HandleFunc("/promotion/{id}", controller.DeletePromotion).Methods(http.MethodDelete)
}
