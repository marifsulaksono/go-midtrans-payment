package errormodel

import (
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/utils/domain"
)

var (
	ErrInternalServer = domain.ErrModel{
		Status:    http.StatusInternalServerError,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   "Please Contact Service",
	}
)
