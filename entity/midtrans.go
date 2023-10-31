package entity

import "github.com/google/uuid"

type MidtransRequestPayload struct {
	PaymentType        string       `json:"payment_type"`
	TransactionDetails OrderDetail  `json:"transaction_details"`
	BankTransfer       BankTransfer `json:"bank_transfer,omitempty"`
	Echannel           BillInfo     `json:"echannel,omitempty"`
}

type OrderDetail struct {
	OrderId  uuid.UUID `json:"order_id"`
	GrossAmt int       `json:"gross_amount"`
}

type BankTransfer struct {
	Bank Banks `json:"bank"`
}
