package models

type AddTransactionRequest struct {
	GiftCode string  `json:"giftCode"`
	Amount   float64 `json:"amount"`
}
