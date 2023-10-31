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
		BankTransfer: entity.BankTransfer{
			Bank: payment.PaymentBank,
		},
		Echannel: payment.Echannel,
	}

	conf := config.GetPaymentConfig()

	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))

	payloadRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, conf.SandboxLink, bytes.NewBuffer(payloadRequest))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+authString)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
