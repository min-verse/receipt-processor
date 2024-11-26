package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
	// "github.com/min-verse/receipt-processor/api"
	// "github.com/min-verse/receipt-processor/internal/tools"
	// log "github.com/sirupsen/logrus"
	// "github.com/gorilla/schema"
)

func CalculateReceiptPoints(w http.ResponseWriter, r *http.Request) {
	var receiptId string = chi.URLParam(r, "id")
	var successMsg string = fmt.Sprintf("Successfully reached receipt with ID %s", receiptId)

	var response map[string]string = map[string]string{"success": successMsg}
	w.Header().Set("Content-Type", "application/json")
	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}
}