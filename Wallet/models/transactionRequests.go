package models

type AddTransactionRequest struct {
	GiftCode *string `json:"giftCode"`
	Amount   int     `json:"amount"`
}
