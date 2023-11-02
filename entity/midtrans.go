package entity

import "github.com/google/uuid"

type MidtransRequestPayload struct {
	PaymentType        Type         `json:"payment_type"`
	TransactionDetails OrderDetail  `json:"transaction_details"`
	BankTransfer       BankTransfer `json:"bank_transfer,omitempty"`
	Echannel           BillInfo     `json:"echannel,omitempty"`
	Store              CStore       `json:"cstore,omitempty"`
}

type OrderDetail struct {
	OrderId  uuid.UUID `json:"order_id"`
	GrossAmt int       `json:"gross_amount"`
}

type BankTransfer struct {
	Bank Banks `json:"bank,omitempty"`
}

type MidtransSnapRequestPayload struct {
	PaymentType        string          `json:"payment_type"`
	TransactionDetails OrderDetail     `json:"transaction_details"`
	CreditCard         Card            `json:"credit_card,omitempty"`
	CustomerDetail     CostumerDetails `json:"customer_details,omitempty"`
}

type Card struct {
	Secure bool `json:"secure,omitempty"`
}

type CostumerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
