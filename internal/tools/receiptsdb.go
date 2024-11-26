package tools

import(
	"time"
)

type receiptsDB struct{}

var receiptsTable = map[string]Receipt{}

func (db *receiptsDB) FindReceipt(receiptId string) *Receipt {
	time.Sleep(time.Second * 0.1)

	var receiptData = Receipt{}
	receiptData, ok := receiptsTable[receiptId]
	if !ok {
		return nil
	}

	return &receiptData
}

func (db *receiptsDB) SetupDatabase() error {
	return nil
}