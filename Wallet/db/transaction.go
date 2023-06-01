package db

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	WalletID uint    `gorm:"uniqueIndex:idx_wallet_id_gift_code"`
	Amount   int     `gorm:"column:amount"`
	GiftCode *string `gorm:"uniqueIndex:idx_wallet_id_gift_code;default:null"`
}
