package main

import (
	"Discount/controllers"
	"Discount/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/gifts/:giftCode", controllers.GetUsersByGiftCode)
	r.POST("/gifts", controllers.CreateGift)
	r.PATCH("/gifts", controllers.UseGift)

	err := r.Run()
	if err != nil {
		return
	}
}
