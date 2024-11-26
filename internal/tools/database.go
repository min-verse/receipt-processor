package tools

import(
	"time"
	// "errors"
	log "github.com/sirupsen/logrus"
)

type Receipt struct{
	Retailer string
	PurchaseDateTime time.Time
	Total float64
	Items []ItemReceipt
}

type ItemReceipt struct{
	ShortDescription string
	Price float64
}

type DatabaseInterface interface{
	// FindReceipt(receiptId string) *Receipt
	FindReceipt(receiptId string) (Receipt, error)
	CreateReceipt(receiptData Receipt) (string, error)
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &receiptsDB{}

	var err error = database.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}