package db

import (
	"gorm.io/gorm"
	"time"
)

type GiftCode struct {
	gorm.Model
	Code           string    `gorm:"uniqueIndex"`
	Description    string    `gorm:"column:description"`
	Amount         float64   `gorm:"column:amount"`
	ExpirationDate time.Time `gorm:"column:expiration_date"`
	MaxUsage       uint      `gorm:"column:max_usage"`
	CurrentUsage   uint      `gorm:"column:current_usage"`
}
