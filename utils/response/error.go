package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/utils/domain"
	errormodel "github.com/marifsulaksono/go-midtrans-payment/utils/domain/errorModel"
)

func ErrorResponseBuilder(w http.ResponseWriter, err error) {
	response := errorBuilder(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func errorBuilder(err error) domain.ErrModel {
	var response domain.ErrModel
	if checkError, ok := err.(domain.ErrModel); ok {
		response.Status = checkError.Status
		response.ErrorCode = checkError.ErrorCode
		response.Message = checkError.Message
		response.Details = checkError.Details
	} else {
		response = errormodel.ErrInternalServer
		log.Printf("Internal Server Error : %v", err)
	}

	return response
}
