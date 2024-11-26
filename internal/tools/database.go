package tools

import(
	"time"
	// "errors"
	log "github.com/sirupsen/logrus"
)

type Receipt struct{
	retailer string
	purchaseDateTime time.Time
	total float64
	items []ItemReceipt
}

type ItemReceipt struct{
	shortDescription string
	price float64
}

type DatabaseInterface interface{
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