package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-midtrans-payment/config"
	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/repository"
)

type PaymentService struct {
	Repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return PaymentService{Repo: repo}
}

func (p *PaymentService) CreatePayment(ctx context.Context, payment *entity.PaymentDetail) (*http.Response, error) {
	payment.OrderID = uuid.New()
	payment.Date = time.Now()
	payment.Status = entity.Waiting

	var paymentRequest entity.MidtransRequestPayload

	// payment type validation
	if payment.PaymentType == entity.Permata || payment.PaymentType == entity.Gopay || payment.PaymentType == entity.Qris || payment.PaymentType == entity.Alku || payment.PaymentType == entity.Kred {
		paymentRequest = entity.MidtransRequestPayload{
			PaymentType: payment.PaymentType,
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: payment.Total,
			},
		}
	} else if payment.PaymentType == entity.Bank {
		paymentRequest = entity.MidtransRequestPayload{
			PaymentType: payment.PaymentType,
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: payment.Total,
			},
			BankTransfer: entity.BankTransfer{
				Bank: payment.PaymentBank,
			},
		}
	} else if payment.PaymentType == entity.Mandiri {
		paymentRequest = entity.MidtransRequestPayload{
			PaymentType: payment.PaymentType,
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: payment.Total,
			},
			Echannel: payment.Echannel,
		}
	} else if payment.PaymentType == entity.Store {
		paymentRequest = entity.MidtransRequestPayload{
			PaymentType: payment.PaymentType,
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: payment.Total,
			},
			Store: payment.Store,
		}
	} else {
		return nil, errors.New("payment type is not allowed")
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

	if response.StatusCode == http.StatusOK {
		err := p.Repo.CreateTransaction(ctx, payment)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (p *PaymentService) CreateSnapPayment(ctx context.Context, payment *entity.PaymentDetail) (*http.Response, error) {
	payment.OrderID = uuid.New()

	paymentRequest := entity.MidtransSnapRequestPayload{
		TransactionDetails: entity.OrderDetail{
			OrderId:  payment.OrderID,
			GrossAmt: payment.Total,
		},
		CreditCard: entity.Card{
			Secure: true,
		},
		CustomerDetail: payment.CustomerDetail,
	}

	payloadRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	conf := config.GetPaymentConfig()
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))

	request, err := http.NewRequest(http.MethodPost, conf.SnapSandboxLink, bytes.NewBuffer(payloadRequest))
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

	if response.StatusCode == http.StatusOK {
		err := p.Repo.CreateTransaction(ctx, payment)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (p *PaymentService) UpdateTransaction(ctx context.Context, id, status string) error {
	var stts entity.Status

	switch status {
	case "capture":
		stts = entity.Success
	case "settlement":
		stts = entity.Success
	case "pending":
		stts = entity.Waiting
	case "deny":
		stts = entity.Cancel
	case "cancel":
		stts = entity.Cancel
	case "failure":
		stts = entity.Cancel
	case "expire":
		stts = entity.Expired
	}

	return p.Repo.UpdateTransaction(ctx, id, stts)
}
