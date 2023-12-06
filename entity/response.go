package entity

type PaymentLinkResponse struct {
	Status     int    `json:"status"`
	OrderId    string `json:"order_id"`
	PaymentUrl string `json:"payment_url"`
}

type PaymentSnapResponse struct {
	Status     int    `json:"status"`
	Token      string `json:"token"`
	PaymentUrl string `json:"payment_url"`
}

type PaymentCoreResponse struct {
	Status    int    `json:"status"`
	PaymentId string `json:"payment_id"`
	VANumber  any    `json:"va_number"`
}
