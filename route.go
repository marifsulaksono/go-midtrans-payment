package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marifsulaksono/go-midtrans-payment/controller"
)

func routeInit() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/payments", controller.CreatePayment).Methods(http.MethodPost)

	return r
}
