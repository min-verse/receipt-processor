package tools

import(
	"fmt"
	"errors"
	"time"
	"github.com/google/uuid"
)

type receiptsDB struct{}

var receiptsTable = map[string]Receipt{}

func (db *receiptsDB) FindReceipt(receiptId string) (Receipt, error) {
	var err error

	time.Sleep(time.Second * 1)

	// var receiptData Receipt
	receiptData, ok := receiptsTable[receiptId]
	if !ok {
		err = errors.New(fmt.Sprintf("Could not find a Receipt with ID %v", receiptId))
		return receiptData, err
	}

	// return &receiptData
	return receiptData, err
}

func (db *receiptsDB) CreateReceipt(newReceiptData Receipt) (string, error) {
	var err error
	// Check to see if id already exists to prevent collisions
	var newReceiptId string = uuid.New().String()

	existingReceiptData, exists := receiptsTable[newReceiptId]
	if exists{
		err = errors.New(fmt.Sprintf("Data already exists for ID %v: %v", newReceiptId, existingReceiptData))
		return newReceiptId, err
	}

	receiptsTable[newReceiptId] = newReceiptData
	var _, ok = receiptsTable[newReceiptId]
	if !ok {
		err = errors.New(fmt.Sprintf("Unable to save ID %v: %v", newReceiptId, newReceiptData))
		return newReceiptId, err
	}

	return newReceiptId, err
}

func (db *receiptsDB) SetupDatabase() error {
	return nil
}