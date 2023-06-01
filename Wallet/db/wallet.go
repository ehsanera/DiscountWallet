package db

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	PhoneNumber  string        `json:"phoneNumber" gorm:"uniqueIndex"`
	Transactions []Transaction `gorm:"foreignKey:WalletID"`
}
