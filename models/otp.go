package models

import (
	"crypto/rand"
	"io"

	"gorm.io/gorm"
)

type Otp struct {
	BaseMode
	Otp      string `json:"otp"`
	IsActive string `json:"isActive"`
}

func (otp *Otp) BeforeCreate(tx *gorm.DB) error {
	otp.Otp = EncodeToString(6)
	return nil
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
