package service

import (
	"Wallet/db"
	"gorm.io/gorm"
)

func CreateTransaction(transaction *db.Transaction) error {
	return db.DB.Create(transaction).Error
}

func FindFirstTransactionByWalletIdAndGiftCode(walletID uint, giftCode *string) (*db.Transaction, error) {
	var existingTransaction db.Transaction
	err := db.DB.Where("wallet_id = ? AND gift_code = ?", walletID, giftCode).First(&existingTransaction).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existingTransaction.ID != 0 {
		return &existingTransaction, nil
	}
	return nil, nil
}
