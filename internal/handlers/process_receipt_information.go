package handlers

import (
	"errors"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/min-verse/receipt-processor/api"
	// "github.com/min-verse/receipt-processor/internal/tools"
	log "github.com/sirupsen/logrus"
	// "github.com/gorilla/schema"
)

var MissingFieldsError = errors.New(`Missing necessary field(s) to create a Receipt record: retailer, purchaseDate, purchaseTime,
and items where each item has a shortDescription and price, both as strings`)

func ProcessReceiptInformation(w http.ResponseWriter, r *http.Request) {
	var receiptPayload api.ReceiptRequest
	
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&receiptPayload)

	if decodeErr != nil{
		log.Error(MissingFieldsError)
		api.RequestErrorHandler(w, MissingFieldsError)
		return
	}

	var itemString string = fmt.Sprintf("%v", receiptPayload.ItemReceipts[0].ShortDescription)

	var response map[string]string = map[string]string{"success": itemString}

	w.Header().Set("Content-Type", "application/json")

	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}

	// var database *tools.DatabaseInterface
	// database, err = tools.NewDatabase()
	// if err != nil {
	// 	api.InternalErrorHandler(w)
	// 	return
	// }

	// fmt.Sprintf("%+v", database)
}