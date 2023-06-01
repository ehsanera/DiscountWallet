package main

import (
	"Wallet/controllers"
	"Wallet/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/wallets/:phoneNumber", controllers.BalanceHandler)
	r.POST("/wallets/:phoneNumber/transactions", controllers.CreateTransactionHandler)
	r.GET("/wallets/getPhoneNumbers/:giftCode", controllers.GetWalletPhoneNumbersByGiftCodeHandler)

	port := ":8081"
	err := r.Run(port)
	if err != nil {
		return
	}
}
