package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	codeCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateDiscountCode(discountCodeLength int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	charsetLen := len(codeCharset)

	for i := 0; i < discountCodeLength; i++ {
		index := rand.Intn(charsetLen)
		sb.WriteByte(codeCharset[index])
	}

	return sb.String()
}
