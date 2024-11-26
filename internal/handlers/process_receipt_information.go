package handlers

import (
	"encoding/json"
	"net/http"
	// "github.com/min-verse/receipt-processor/api"
	// "github.com/min-verse/receipt-processor/internal/tools"
	// log "github.com/sirupsen/logrus"
	// "github.com/gorilla/schema"
)

func ProcessReceiptInformation(w http.ResponseWriter, r *http.Request) {
	var response map[string]string = map[string]string{"success":"you've reached here"}
	w.Header().Set("Content-Type", "application/json")
	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}
}