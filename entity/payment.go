package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Banks  string
	Status string
	Type   string
)

const (
	BCA  Banks = "bca"
	BRI  Banks = "bri"
	BNI  Banks = "bni"
	CIMB Banks = "cimb"

	Bank    Type = "bank_transfer"
	Mandiri Type = "echannel"
	Permata Type = "permata"
	Store   Type = "cstore"
	Alku    Type = "akulaku"
	Kred    Type = "kredivo"
	Gopay   Type = "gopay"
	Qris    Type = "qris"

	Waiting Status = "waiting"
	Cancel  Status = "cancel"
	Expired Status = "expired"
	Success Status = "success"
)

type PaymentDetail struct {
	OrderID        uuid.UUID       `json:"order_id"`
	Date           time.Time       `json:"date"`
	Total          *int            `json:"total"`
	Status         Status          `json:"status"`
	PaymentType    Type            `json:"payment_type"`
	PaymentBank    Banks           `json:"payment_bank,omitempty"`
	Echannel       BillInfo        `json:"echannel,omitempty"`
	Store          CStore          `json:"store,omitempty"`
	CustomerDetail CostumerDetails `json:"customer_details,omitempty"`
	ItemDetail     []ItemDetails   `json:"item_details"`
}

type BillInfo struct {
	BillInfo1 string `json:"bill_info1,omitempty"`
	BillInfo2 string `json:"bill_info2,omitempty"`
}

type CStore struct {
	Store   string `json:"store,omitempty"`
	Message string `json:"message,omitempty"`
}

type ItemDetails struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Qty          int    `json:"quantity"`
	MerchantName string `json:"merchant_name,omitempty"`
}
