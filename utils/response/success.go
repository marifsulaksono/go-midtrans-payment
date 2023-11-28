package response

import (
	"encoding/json"
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/utils/domain"
)

func SuccessResponseBuilder(w http.ResponseWriter, statusCode int, data, metadata interface{}, message string) {
	payload := domain.SuccessModel{
		Status:   statusCode,
		Data:     data,
		Metadata: metadata,
		Message:  message,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
