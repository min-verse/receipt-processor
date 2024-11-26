package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux){
	// Standardizes the request and strips slashes
	r.Use(chimiddle.StripSlashes)

	// Defines the "/receipts" namespace for routes
	r.Route("/receipts", func(router chi.Router){
		// Handles POST requests for new Receipts
		router.Post("/process", ProcessReceiptInformation)

		// Handles GET requests for existing Receipt's total points
		router.Get("/{id}/points", CalculateReceiptPoints)
	})
}