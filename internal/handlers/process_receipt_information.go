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

	// var itemString string = fmt.Sprintf("%v is of type %T", receiptPayload.ItemReceipts[0].Price, receiptPayload.ItemReceipts[0].Price)

	combinedDateTime := fmt.Sprintf("%sT%s:00", receiptPayload.PurchaseDate, receiptPayload.PurchaseTime)
	parsedTime, dateErr := time.Parse("2006-01-02T15:04:05", combinedDateTime)
	if dateErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	var itemStructSlice []tools.ItemReceipt
	for i := range receiptPayload.ItemReceipts{
		currItem := receiptPayload.ItemReceipts[i]
		var newItem tools.ItemReceipt = tools.ItemReceipt{ShortDescription: currItem.ShortDescription, Price: currItem.Price}
		itemStructSlice = append(itemStructSlice, newItem)
	}

	var receiptInstance tools.Receipt = tools.Receipt{Retailer: receiptPayload.Retailer, PurchaseDateTime: parsedTime, Total: receiptPayload.Total, Items: itemStructSlice}
	// var newReceiptId string = fmt.Sprintf("%d", rand.Intn(100)+10)

	var database *tools.DatabaseInterface
	database, databaseErr := tools.NewDatabase()
	if databaseErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	savedReceiptId, saveErr := (*database).CreateReceipt(receiptInstance)
	if saveErr != nil {
		api.InternalErrorHandler(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	
	// var response map[string]string = map[string]string{"success": fmt.Sprintf("%v are the items of type %T, and the date is %v converted to datetime as %v which is of type %T", receiptInstance.Items, receiptInstance.Items, combinedDateTime, receiptInstance.PurchaseDateTime, receiptInstance.PurchaseDateTime)}
	var response map[string]string = map[string]string{"id": savedReceiptId}
	var err = json.NewEncoder(w).Encode(response)
	if err != nil{
		return
	}


	// fmt.Sprintf("%+v", database)
}