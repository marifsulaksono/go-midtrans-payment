package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marifsulaksono/go-midtrans-payment/controller"
)

func routeInit() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/payments/core", controller.CreatePayment).Methods(http.MethodPost)
	r.HandleFunc("/payments/snap", controller.CreateSnapPayment).Methods(http.MethodPost)
	r.HandleFunc("/payments/notification", controller.WebhookPayment).Methods(http.MethodPost)

	return r
}
