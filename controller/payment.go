package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/service"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var payment entity.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.CreatePayment(ctx, &payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "IO Reader Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseJSON)
}

func CreateSnapPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var payment entity.PaymentDetail

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.CreateSnapPayment(ctx, &payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "IO Read Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseJSON)
}
