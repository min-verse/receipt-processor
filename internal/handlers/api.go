package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	// Imports middleware package
	"github.com/min-verse/receipt-processor/internal/middleware"
)

func Handler(r *chi.Mux){
	// Standardizes the request and strips slashes
	r.Use(chimiddle.StripSlashes)

	r.Route("/receipts", func(router chi.Router){
		router.Post("/process", ProcessReceiptInformation)

		router.Get("/{id}/points", CalculateReceiptPoints)
	})
}