package service

import (
	"errors"
	"time"

	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/utils/helper"
)

func PayloadRequestMidtransBuilder(method string, payment *entity.PaymentDetail) (entity.MidtransPayloadRequest, error) {
	var paymentRequest entity.MidtransPayloadRequest

	// payment validation
	if method == "link" {
		var (
			cr    = true
			usage = 1
		)

		paymentRequest = entity.MidtransPayloadRequest{
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: *payment.Total,
			},
			CustomerRequired: &cr,
			Usage:            &usage,
			Expiry: &entity.ExpiryDetails{
				Start:    time.Now().Format("2006-01-02 15:04:05 +0700"),
				Duration: 24,
				Unit:     "hours",
			},
			ItemDetail: payment.ItemDetail,
			CustomerDetail: &entity.CostumerDetails{
				FirstName: payment.CustomerDetail.FirstName,
				LastName:  payment.CustomerDetail.LastName,
				Email:     payment.CustomerDetail.Email,
				Phone:     payment.CustomerDetail.Phone,
			},
		}
	} else if method == "snap" {
		paymentRequest = entity.MidtransPayloadRequest{
			TransactionDetails: entity.OrderDetail{
				OrderId:  payment.OrderID,
				GrossAmt: *payment.Total,
			},
			Expiry: &entity.ExpiryDetails{
				Start:    time.Now().Format("2006-01-02 15:04:05 +0700"),
				Duration: 24,
				Unit:     "hours",
			},
			ItemDetail: payment.ItemDetail,
			CustomerDetail: &entity.CostumerDetails{
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
					GrossAmt: *payment.Total,
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
			}
		} else {
			return entity.MidtransPayloadRequest{}, errors.New("payment type is not allowed")
		}
	} else {
		return entity.MidtransPayloadRequest{}, errors.New("payment method is undefined")
	}

	return paymentRequest, nil
}
