package service

import (
	"Wallet/db"
	"gorm.io/gorm"
)

func FindOrCreateWalletByPhoneNumber(phoneNumber string) (*db.Wallet, error) {
	var wallet db.Wallet
	err := db.DB.FirstOrCreate(&wallet, db.Wallet{PhoneNumber: phoneNumber}).Error
	return &wallet, err
}

func GetWalletPhoneNumbersByGiftCode(giftCode string) ([]string, error) {
	var phoneNumbers []string

	// Query to retrieve distinct phone numbers from wallets based on gift code
	err := db.DB.Table("wallets").
		Select("DISTINCT wallets.phone_number").
		Joins("JOIN transactions ON wallets.id = transactions.wallet_id").
		Where("transactions.gift_code = ?", giftCode).
		Pluck("wallets.phone_number", &phoneNumbers).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return phoneNumbers, nil
}
