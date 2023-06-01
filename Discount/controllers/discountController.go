package controllers

import (
	"Discount/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"Discount/db"
	"Discount/models"
	"Discount/service"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
	"gorm.io/gorm"
)

var cb *gobreaker.CircuitBreaker

func init() {
	cb = gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:        "my-circuit-breaker",
			MaxRequests: 3,
			Timeout:     3 * time.Second,
			Interval:    1 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > 3
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		},
	)
}

var (
	wg  sync.WaitGroup
	mux sync.Mutex
)

func UseGift(c *gin.Context) {
	var req models.DiscountUseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !util.ValidatePhoneNumber(req.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	wg.Add(1)
	defer wg.Done()
	mux.Lock()
	defer mux.Unlock()

	giftCode, err := service.FindFirstGiftCodeByCode(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gift code not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the gift code"})
		}
		return
	}

	if giftCode.CurrentUsage == giftCode.MaxUsage {
		c.JSON(http.StatusForbidden, gin.H{"error": "This code is no longer valid"})
		return
	}

	println(req.PhoneNumber)
	err = service.UpdateGiftCodeUsage(giftCode, giftCode.CurrentUsage+1)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Concurrent update conflict"})
		return
	}

	resultI, err := cb.Execute(func() (interface{}, error) {
		result, err := service.CallAnotherAPI("wallets/"+req.PhoneNumber+"/transactions", "POST", models.AddTransactionRequest{
			GiftCode: giftCode.Code,
			Amount:   giftCode.Amount,
		})
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, ok := resultI.(models.ApiResponse)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid API response"})
		return
	}

	statusCode := response.StatusCode
	if statusCode == http.StatusConflict {
		c.JSON(http.StatusConflict, gin.H{"error": "You can't use a gift code multiple times"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "SMS sent"})
}

func CreateGift(c *gin.Context) {
	var req models.DiscountCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := generateDiscountCode(req.Code)

	discount := db.GiftCode{
		Code:     code,
		MaxUsage: req.MaxUse,
		Amount:   req.Amount,
	}

	if err := db.DB.Create(&discount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error create code": err.Error(), "code": code})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": discount})
}

func generateDiscountCode(code string) string {
	if len(code) < 1 {
		return util.GenerateDiscountCode(8)
	}
	return code
}

func GetUsersByGiftCode(c *gin.Context) {
	giftCode := c.Param("giftCode")

	result, err := cb.Execute(func() (interface{}, error) {
		return service.CallAnotherAPI("wallets/getPhoneNumbers/"+giftCode, "GET", nil)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, ok := result.(models.ApiResponse)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid API response"})
		return
	}

	var phoneNumbers []string
	err = json.Unmarshal(response.Bytes, &phoneNumbers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse phone numbers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"phoneNumbers": phoneNumbers})
}
