package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/min-verse/receipt-processor/api"
	"github.com/min-verse/receipt-processor/internal/tools"
	log "github.com/sirupsen/logrus"
)

func CalculateReceiptPoints(w http.ResponseWriter, r *http.Request) {
	// Retrieves the ID from the URL
	var receiptId string = chi.URLParam(r, "id")

	// Simulating setting up a database connection
	// (Built-in 1 second delay to simulate latency)
	var database *tools.DatabaseInterface
	database, databaseErr := tools.NewDatabase()
	if databaseErr != nil {
		log.Error(databaseErr)
		api.InternalErrorHandler(w)
		return
	}

	// Simulating an ORM call to find a Receipt record
	// and errors out if there is no such record found
	receipt, retrievalErr := (*database).FindReceipt(receiptId)
	if retrievalErr != nil{
		log.Error(retrievalErr)
		api.RequestErrorHandler(w, retrievalErr)
		return
	}

	// Method defined on Receipt Struct to get total points
	var totalPoints int = receipt.CalculateTotalPoints()

	// Defines the response with total points and a 200 status code by default
	var response map[string]int = map[string]int{"points": totalPoints}
	w.Header().Set("Content-Type", "application/json")
	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}
}