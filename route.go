package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marifsulaksono/go-midtrans-payment/controller"
	"github.com/marifsulaksono/go-midtrans-payment/repository"
	"github.com/marifsulaksono/go-midtrans-payment/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func routeInit(conn *mongo.Client) *mux.Router {
	paymentRepo := repository.NewPaymentRepository(conn)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)

	r := mux.NewRouter()

	r.HandleFunc("/payments", paymentController.CreateNewPayment).Methods(http.MethodPost)
	r.HandleFunc("/payments/notification", paymentController.WebhookPayment).Methods(http.MethodPost)

	return r
}
