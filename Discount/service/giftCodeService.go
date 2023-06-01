package service

import (
	"Discount/db"
)

func FindFirstGiftCodeByCode(code string) (db.GiftCode, error) {
	var giftCode db.GiftCode
	db.DB.Where("code = ?", code).First(&giftCode)
	return giftCode, nil
}

func UpdateGiftCodeUsage(giftCode db.GiftCode, currentUsage uint) error {
	giftCode.CurrentUsage = currentUsage
	return db.DB.Save(&giftCode).Error
}
