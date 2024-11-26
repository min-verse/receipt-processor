package tools

import(
	"fmt"
	"errors"
	"time"
	"github.com/google/uuid"
)

// The "database" struct
type receiptsDB struct{}

// The receipts "table" in the "database"
var receiptsTable = map[string]Receipt{}

// Retrieves receiptId from the receipts "table"
// This takes a string, receiptId, and then returns a Receipt Struct
// and also returns an error if something went wrong
func (db *receiptsDB) FindReceipt(receiptId string) (Receipt, error) {
	var err error

	// Manually added this in to simulate database latency
	time.Sleep(time.Second * 1)

	// Checks to see if data even exists in the "database"
	receiptData, ok := receiptsTable[receiptId]
	if !ok {
		err = errors.New(fmt.Sprintf("Could not find a Receipt with ID %v", receiptId))
		return receiptData, err
	}

	// returns the Receipt Struct found and any errors
	return receiptData, err
}

// Saves Receipt Struct to the "database"
// This takes a Receipt Struct and returns the string UUID
// and also returns an error if something went wrong
func (db *receiptsDB) CreateReceipt(newReceiptData Receipt) (string, error) {
	var err error

	var newReceiptId string = uuid.New().String()
	
	// Check to see if UUID already exists to prevent collisions
	// Essentially I am manually writing checks that are usually
	// under the hood in most ORM technologies
	existingReceiptData, exists := receiptsTable[newReceiptId]
	if exists{
		err = errors.New(fmt.Sprintf("Data already exists for ID %v: %v", newReceiptId, existingReceiptData))
		return newReceiptId, err
	}

	// Saves Receipt Struct to "database"
	receiptsTable[newReceiptId] = newReceiptData
	var _, ok = receiptsTable[newReceiptId]
	if !ok {
		err = errors.New(fmt.Sprintf("Unable to save ID %v: %v", newReceiptId, newReceiptData))
		return newReceiptId, err
	}

	// Returns the UUID for this Receipt Struct record and any errors
	return newReceiptId, err
}

// Simulates "setting up" a "database"
func (db *receiptsDB) SetupDatabase() error {
	return nil
}