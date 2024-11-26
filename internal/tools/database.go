package tools

import(
	"unicode"
	"math"
	"strings"
	"time"
	log "github.com/sirupsen/logrus"
)

// Data Models are saved here
// Can think of this a simulated (NoSQL) "database"

// Receipt Struct
type Receipt struct{
	Retailer string
	PurchaseDateTime time.Time
	Total float64
	Items []ItemReceipt
}

// Related Receipt Struct
// One Receipt Struct "has many" ItemReceipt Structs
type ItemReceipt struct{
	ShortDescription string
	Price float64
}

// Database Interface with Method Signatures such that
// it simulates using something like an ORM
type DatabaseInterface interface{
	FindReceipt(receiptId string) (Receipt, error)
	CreateReceipt(receiptData Receipt) (string, error)
	SetupDatabase() error
}

// Creates the "database"
func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &receiptsDB{}

	var err error = database.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

// Method to call on Receipt Struct to calculate total points
func (receipt Receipt) CalculateTotalPoints() int {
	// Calculating Points
	var totalPoints int = 0

	// alphanumeric runes: +1 each
	for _, r := range receipt.Retailer{
		if unicode.IsLetter(r) || unicode.IsNumber(r){
			totalPoints++
		}
	}

	// The Total number is whole number: +50
	if receipt.Total == math.Trunc(receipt.Total){
		totalPoints+=50
	}

	// The Total number is a multiple of 0.25: +25
	if math.Mod(receipt.Total, 0.25) == 0{
		totalPoints+=25
	}
	
	// For every 2 items on the Receipt Struct: +5 each
	var quotient int = len(receipt.Items) / 2
	var itemPoints int = quotient * 5
	totalPoints+=itemPoints

	// For each ItemReceipt Struct's shortDescription,
	// IF its trimmed length a multiple of 3,
	// multiply the ItemReceipt Struct's price by 0.2,
	// and then round UP to nearest integer,
	// then add to totalPoints
	for _, v := range receipt.Items{
		trimmedDesc := strings.TrimSpace(v.ShortDescription)
		if len(trimmedDesc)%3 == 0{
			priceProduct := v.Price * 0.2

			totalPoints+=int(math.Ceil(priceProduct))
		}
	}

	// If the PurchaseDateTime is odd: +6
	if receipt.PurchaseDateTime.Day()%2 != 0{
		totalPoints+=6
	}

	// If the PurchaseDateTime is between 2PM and 4PM,
	// assuming that the range is exclusive
	// (we do not accept 2PM exactly or 4PM exactly)
	var purchaseHour int = receipt.PurchaseDateTime.Hour()
	if purchaseHour >= 14 && purchaseHour < 16{
		// Because we're assuming 2:00PM exactly is
		// not included, we must see if the minute value
		// is not equal to 0 and therefore greater
		if purchaseHour == 14{
			if receipt.PurchaseDateTime.Minute() > 0{
				totalPoints+=10
			}
		}else{
			totalPoints+=10
		}
	}

	return totalPoints
}