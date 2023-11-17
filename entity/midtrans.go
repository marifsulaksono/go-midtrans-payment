package entity

import (
	"github.com/google/uuid"
)

type MidtransPayloadRequest struct {
	PaymentType        Type            `json:"payment_type"`
	TransactionDetails OrderDetail     `json:"transaction_details"`
	BankTransfer       BankTransfer    `json:"bank_transfer,omitempty"`
	Echannel           BillInfo        `json:"echannel,omitempty"`
	Store              CStore          `json:"cstore,omitempty"`
	CustomerRequired   *bool           `json:"customer_required,omitempty"`
	Usage              *int            `json:"usage_limit,omitempty"`
	Expiry             *ExpiryDetails  `json:"expiry,omitempty"`
	ItemDetail         []ItemDetails   `json:"item_details,omitempty"`
	CustomerDetail     CostumerDetails `json:"customer_details,omitempty"`
}

type OrderDetail struct {
	OrderId  uuid.UUID `json:"order_id"`
	GrossAmt int       `json:"gross_amount"`
	// PaymentLinkId string   `json:"payment_link_id,omitempty"`
}

type BankTransfer struct {
	Bank Banks `json:"bank,omitempty"`
}

type CostumerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ExpiryDetails struct {
	Start    string `json:"start_time"`
	Duration int    `json:"duration"`
	Unit     string `json:"unit"`
}
