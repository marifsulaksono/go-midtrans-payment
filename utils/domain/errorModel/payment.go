package errormodel

import (
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/utils/domain"
)

var (
	ErrCreatePayment = domain.ErrModel{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_PAYMENT_INPUT",
		Message:   "Some required input is missing",
	}
)
