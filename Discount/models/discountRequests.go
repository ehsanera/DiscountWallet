package models

type DiscountCreateRequest struct {
	Code   string  `json:"code"`
	MaxUse uint    `json:"maxUse"`
	Amount float64 `json:"amount"`
}

type DiscountUseRequest struct {
	Code        string `json:"code"`
	PhoneNumber string `json:"phoneNumber"`
}
