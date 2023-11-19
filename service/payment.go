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
	"github.com/marifsulaksono/go-midtrans-payment/utils/helper"
)

type PaymentService struct {
	Repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return PaymentService{Repo: repo}
}

func (p *PaymentService) CreateNewPayment(ctx context.Context, method string, payment *entity.PaymentDetail) (*http.Response, error) {
	payment.OrderID = uuid.New()
	payment.Date = time.Now()
	payment.Status = entity.Waiting

	var (
		paymentRequest entity.MidtransPayloadRequest
		cr             = true
		usage          = 1
	)

	// payment validation
	if method == "link" || method == "snap" {
		paymentRequest = entity.MidtransPayloadRequest{
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: payment.Total,
			},
			CustomerRequired: &cr,
			Usage:            &usage,
			Expiry: &entity.ExpiryDetails{
				Start:    time.Now().Format("2006-01-02 15:04:05 +0700"),
				Duration: 24,
				Unit:     "hours",
			},
			ItemDetail: payment.ItemDetail,
			CustomerDetail: entity.CostumerDetails{
				FirstName: payment.CustomerDetail.FirstName,
				LastName:  payment.CustomerDetail.LastName,
				Email:     payment.CustomerDetail.Email,
				Phone:     payment.CustomerDetail.Phone,
			},
		}
	} else if method == "core" {
		if helper.IsValidPaymentType(payment.PaymentType) {
			paymentRequest = entity.MidtransPayloadRequest{
				PaymentType: payment.PaymentType,
				TransactionDetails: entity.OrderDetail{
					OrderId:  payment.OrderID,
					GrossAmt: payment.Total,
				},
				BankTransfer: entity.BankTransfer{
					Bank: payment.PaymentBank,
				},
				Echannel: entity.BillInfo{
					BillInfo1: payment.Echannel.BillInfo1,
					BillInfo2: payment.Echannel.BillInfo2,
				},
				Store: entity.CStore{
					Store:   payment.Store.Store,
					Message: payment.Store.Message,
				},
				ItemDetail: payment.ItemDetail,
				CustomerDetail: entity.CostumerDetails{
					FirstName: payment.CustomerDetail.FirstName,
					LastName:  payment.CustomerDetail.LastName,
					Email:     payment.CustomerDetail.Email,
					Phone:     payment.CustomerDetail.Phone,
				},
			}
		} else {
			return nil, errors.New("payment type is not allowed")
		}
	} else {
		return nil, errors.New("payment method is undefined")
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
	var link string
	if method == "core" {
		link = conf.CoreSandboxLink
	} else if method == "snap" {
		link = conf.SnapSandboxLink
	} else if method == "link" {
		link = conf.SandboxLink
	} else {
		return nil, errors.New("unknown method")
	}

	request, err := http.NewRequest(http.MethodPost, link, bytes.NewBuffer(payloadRequest))
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

	if response.StatusCode == http.StatusCreated {
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
