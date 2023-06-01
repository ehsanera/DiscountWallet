package controllers

import (
	"net/http"

	"Wallet/db"
	"Wallet/models"
	"Wallet/service"
	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(c *gin.Context) {
	var req models.AddTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phoneNumber := c.Param("phoneNumber")
	wallet, err := service.FindOrCreateWalletByPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or retrieve wallet"})
		return
	}

	transaction := db.Transaction{
		WalletID: wallet.ID,
		Amount:   req.Amount,
		GiftCode: req.GiftCode,
	}

	existingTransaction, err := service.FindFirstTransactionByWalletIdAndGiftCode(transaction.WalletID, transaction.GiftCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve existing transaction"})
		return
	}

	if existingTransaction != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Transaction with the same WalletID and GiftCode already exists"})
		return
	}

	if err := service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}
