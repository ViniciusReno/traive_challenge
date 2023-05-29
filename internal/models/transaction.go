package models

import (
	"errors"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type OperationType string

const (
	Credit OperationType = "credit"
	Debit  OperationType = "debit"
	Pix    OperationType = "pix"
)

var operationTypeStrings = map[OperationType]string{
	Credit: "Credit",
	Debit:  "Debit",
	Pix:    "Pix",
}

type Origin string

const (
	DesktopWeb    Origin = "desktop-web"
	MobileAndroid Origin = "mobile-android"
	MobileIOS     Origin = "mobile-ios"
)

var originStrings = map[Origin]string{
	DesktopWeb:    "Desktop Web",
	MobileAndroid: "Mobile Android",
	MobileIOS:     "Mobile iOS",
}

func ParseOrigin(s string) error {
	s = strings.ToLower(s)
	_, ok := originStrings[Origin(s)]
	if !ok {
		return errors.New("invalid origin")
	}
	return nil
}

func ParseOperationType(s string) error {
	s = strings.ToLower(s)
	_, ok := operationTypeStrings[OperationType(s)]
	if !ok {
		return errors.New("invalid operation type")
	}
	return nil
}

type Transaction struct {
	Base

	UserID        uuid.UUID     `gorm:"not null"`
	User          User          `gorm:"foreignkey:UserID"`
	Origin        Origin        `gorm:"not null"`
	Amount        float64       `gorm:"not null"`
	OperationType OperationType `gorm:"not null"`
}

type TransactionResponse struct {
	Amount        float64       `json:"amount"`
	CreatedAt     time.Time     `json:"create_at"`
	OperationType OperationType `json:"opeation_type"`
	Origin        Origin        `json:"origin"`
	TransactionID uuid.UUID     `json:"transaction_id"`
	UpdatedAt     time.Time     `json:"update_at"`
	UserID        uuid.UUID     `json:"user_id"`
}
