package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Banks  string
	Status string
)

const (
	BCA  Banks = "bca"
	BRI  Banks = "bri"
	BNI  Banks = "bni"
	CIMB Banks = "cimb"

	Waiting Status = "waiting"
	Cancel  Status = "cancel"
	Expired Status = "expired"
	Success Status = "success"
)

type PaymentDetail struct {
	OrderID     uuid.UUID `json:"order_id"`
	Date        time.Time `json:"date"`
	Total       int       `json:"total"`
	Status      Status    `json:"status"`
	PaymentType string    `json:"payment_type"`
	PaymentBank Banks     `json:"payment_bank,omitempty"`
	Echannel    BillInfo  `json:"ehannel,omitempty"`
}

type BillInfo struct {
	BillInfo1 string `json:"bill_info1"`
	BillInfo2 string `json:"bill_info2"`
}
