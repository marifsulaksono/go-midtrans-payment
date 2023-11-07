package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/logger"
	"github.com/marifsulaksono/go-midtrans-payment/service"
)

type PaymentController struct {
	Service service.PaymentService
}

func NewPaymentController(s service.PaymentService) PaymentController {
	return PaymentController{Service: s}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	// open file logger
	logger, err := logger.OpenFileErrorLogger("./logger/logger.log")
	if err != nil {
		http.Error(w, "Error open log file : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer logger.Close()

	var payment entity.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		log.Println("JSON Core Payment Error : " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := p.Service.CreatePayment(ctx, &payment)
	if err != nil {
		log.Printf("Error Create Core Payment : %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("IO Core Reader Error : %v", err.Error())
		http.Error(w, "IO Core Reader Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseJSON)
}

func (p *PaymentController) CreateSnapPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var payment entity.PaymentDetail

	// open file logger
	logger, err := logger.OpenFileErrorLogger("./logger/logger.log")
	if err != nil {
		http.Error(w, "Error open log file : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer logger.Close()

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		log.Printf("JSON Snap Payment Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := p.Service.CreateSnapPayment(ctx, &payment)
	if err != nil {
		log.Printf("Error Create Snap Payment : %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("IO Snap Reader Error : %v", err.Error())
		http.Error(w, "IO Snap Reader Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseJSON)
}

func (p *PaymentController) WebhookPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	// open file logger notifcation
	logger, err := logger.OpenFileErrorLogger("./logger/logger.log")
	if err != nil {
		http.Error(w, "Error open log file : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer logger.Close()

	var notification map[string]any
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		log.Printf("JSON Webhook Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// logger new notification webhook
	log.Printf("New notification incoming : %v", notification)

	id := fmt.Sprint(notification["order_id"])
	status := fmt.Sprint(notification["transaction_status"])

	err = p.Service.UpdateTransaction(ctx, id, status)
	if err != nil {
		log.Printf("Error Update Incoming Transaction : %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success update transaction"))
}
