package models

type DiscountResponse struct {
	ID     uint `json:"id"`
	Code   int  `json:"code"`
	MaxUse uint `json:"maxUse"`
	Amount uint `json:"amount"`
}
