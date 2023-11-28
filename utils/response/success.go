package response

import (
	"encoding/json"
	"net/http"
)

type successModel struct {
	Status   int         `json:"status_code"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Message  string      `json:"message,omitempty"`
}

func SuccessResponseBuilder(w http.ResponseWriter, statusCode int, data, metadata interface{}, message string) {
	payload := successModel{
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
