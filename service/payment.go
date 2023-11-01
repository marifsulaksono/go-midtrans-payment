package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-midtrans-payment/config"
	"github.com/marifsulaksono/go-midtrans-payment/entity"
)

func CreatePayment(ctx context.Context, payment *entity.PaymentDetail) (*http.Response, error) {
	payment.OrderID = uuid.New()
	payment.Date = time.Now()
	payment.Status = entity.Waiting

	paymentRequest := entity.MidtransRequestPayload{
		PaymentType: payment.PaymentType,
		TransactionDetails: entity.OrderDetail{
			OrderId:  payment.OrderID,
			GrossAmt: payment.Total,
		},
		// this property will not display is payment type isn't bank transfer (bca/bri/bni/cimb)
		BankTransfer: entity.BankTransfer{
			Bank: payment.PaymentBank,
		},
		// this property will not display is payment type isn't echannel (Mandiri bill)
		Echannel: payment.Echannel,
		// this property will not display is payment type isn't over the counter
		Store: payment.Store,
	}

	payloadRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	// get the midtrans configuration
	conf := config.GetPaymentConfig()

	// encode server key to base64
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))

	// request set-up
	request, err := http.NewRequest(http.MethodPost, conf.SandboxLink, bytes.NewBuffer(payloadRequest))
	if err != nil {
		return nil, err
	}

	// header set-up
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+authString)

	// hit midtrans API enpoint with the prepared request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func CreateSnapPayment(ctx context.Context, payment *entity.PaymentDetail) (*http.Response, error) {
	payment.OrderID = uuid.New()

	paymentRequest := entity.MidtransRequestPayload{
		TransactionDetails: entity.OrderDetail{
			OrderId:  payment.OrderID,
			GrossAmt: payment.Total,
		},
		CreditCard: entity.Card{
			Secure: true,
		},
		CustomerDetail: payment.CustomerDetail,
	}

	payloadJSON, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	conf := config.GetPaymentConfig()
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))

	request, err := http.NewRequest(http.MethodPost, conf.SnapSandboxLink, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+authString)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
