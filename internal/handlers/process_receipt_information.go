package handlers

import (
	"time"
	"errors"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/min-verse/receipt-processor/api"
	"github.com/min-verse/receipt-processor/internal/tools"
	log "github.com/sirupsen/logrus"
)

// Custom error to inform the end user of missing fields when attempting
// to create a Receipt Struct to add to our "database"

// Soft enforcement of fields needed
var MissingFieldsError = errors.New(`Missing necessary field(s) to create a Receipt record: retailer, purchaseDate, purchaseTime,
and items where each item has a shortDescription and price, both as strings`)

func ProcessReceiptInformation(w http.ResponseWriter, r *http.Request) {
	// Initializes receiptPayload to represent request body
	// Applying strong parameters
	var receiptPayload api.ReceiptRequest
	
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&receiptPayload)

	if decodeErr != nil{
		log.Error(MissingFieldsError)
		api.RequestErrorHandler(w, MissingFieldsError)
		return
	}

	// Combining PurchaseDate and PurchaseTime to save as time.Time
	// Down the line, we'll want to use this to parse the day and hour
	combinedDateTime := fmt.Sprintf("%sT%s:00", receiptPayload.PurchaseDate, receiptPayload.PurchaseTime)
	parsedTime, dateErr := time.Parse("2006-01-02T15:04:05", combinedDateTime)
	if dateErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Constructs related slice of ItemReceipt Struct
	var itemStructSlice []tools.ItemReceipt
	for i := range receiptPayload.ItemReceipts{
		currItem := receiptPayload.ItemReceipts[i]
		var newItem tools.ItemReceipt = tools.ItemReceipt{ShortDescription: currItem.ShortDescription, Price: currItem.Price}
		itemStructSlice = append(itemStructSlice, newItem)
	}

	// Instantiates a Receipt struct
	var receiptInstance tools.Receipt = tools.Receipt{Retailer: receiptPayload.Retailer, PurchaseDateTime: parsedTime, Total: receiptPayload.Total, Items: itemStructSlice}

	// Just like in CalculateReceiptPoints, simulates creating
	// a "database" connection to eventually save a Receipt to
	var database *tools.DatabaseInterface
	database, databaseErr := tools.NewDatabase()
	if databaseErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Saves the Receipt Struct to the "database" and also captures
	// if there was an error. If there is an error, log it
	savedReceiptId, saveErr := (*database).CreateReceipt(receiptInstance)
	if saveErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Send back a 201 Status Code along with the UUID of the newly-saved Receipt Struct
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Structure of the JSON response back to the end user
	var response map[string]string = map[string]string{"id": savedReceiptId}
	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}
}