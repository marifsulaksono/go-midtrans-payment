package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/service"
	paymentError "github.com/marifsulaksono/go-midtrans-payment/utils/domain/errorModel"
	"github.com/marifsulaksono/go-midtrans-payment/utils/logger"
	buildResponse "github.com/marifsulaksono/go-midtrans-payment/utils/response"
)

type PaymentController struct {
	Service service.PaymentService
}

func NewPaymentController(s service.PaymentService) PaymentController {
	return PaymentController{Service: s}
}

func (p *PaymentController) CreateNewPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m := r.URL.Query().Get("m")

	var payment entity.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		log.Println("JSON Core Payment Error : " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		detailError = make(map[string]any)
	)

	if payment.Total == nil {
		detailError["total"] = "this field is missing input"
	}

	if payment.PaymentType == "" {
		detailError["payment_type"] = "this field is missing input"
	}

	if len(detailError) > 0 {
		err := paymentError.ErrCreatePayment.AttachDetail(detailError)
		buildResponse.ErrorResponseBuilder(w, err)
		return
	}

	response, err := p.Service.CreateNewPayment(ctx, m, &payment)
	if err != nil {
		log.Printf("Error Create Payment : %v", err.Error())
		buildResponse.ErrorResponseBuilder(w, err)
		return
	}

	var responseJSON interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		log.Printf("IO Payment Reader Error : %v", err.Error())
		http.Error(w, "IO Payment Reader Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("New Payment Created : %v", responseJSON)
	buildResponse.SuccessResponseBuilder(w, 201, responseJSON, nil, "Success Create New Payment")
}

func (p *PaymentController) WebhookPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	// open file logger notifcation
	logger, err := logger.OpenFileLogger("./utils/logger/logger.log")
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

	// log new notification webhook
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
