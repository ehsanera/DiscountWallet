package controllers

import (
	"Wallet/db"
	"Wallet/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BalanceHandler(c *gin.Context) {
	var totalAmount int64
	phoneNumber := c.Param("phoneNumber")

	// Execute the SQL query to calculate the sum of transaction amounts
	err := db.DB.Model(&db.Transaction{}).
		Joins("JOIN wallets ON transactions.wallet_id = wallets.id").
		Where("wallets.phone_number = ?", phoneNumber).
		Select("SUM(transactions.amount)").
		Scan(&totalAmount).
		Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalAmount": totalAmount})
}

func GetWalletPhoneNumbersByGiftCodeHandler(c *gin.Context) {
	giftCode := c.Param("giftCode")
	phoneNumbers, err := service.GetWalletPhoneNumbersByGiftCode(giftCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phoneNumbers)
}
