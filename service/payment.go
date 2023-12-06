package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"github.com/marifsulaksono/go-midtrans-payment/repository"
)

type PaymentService struct {
	Repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return PaymentService{Repo: repo}
}

func (p *PaymentService) CreateLinkPaymentMidtrans(ctx context.Context, payment *entity.PaymentDetail) (entity.PaymentLinkResponse, error) {
	// UserId := ctx.Value("user_id").(int)
	UserId := 1
	payment.OrderID = uuid.New()

	paymentRequest, err := PayloadRequestMidtransBuilder("link", payment)
	if err != nil {
		return entity.PaymentLinkResponse{}, err
	}

	response, err := p.Repo.CreateLinkPaymentMidtrans(ctx, &paymentRequest)
	if err != nil {
		return entity.PaymentLinkResponse{}, err
	}

	// insert payment data if create midtrans transaction success
	if response.Status == http.StatusOK || response.Status == http.StatusCreated {
		log.Printf("New Link Payment Created : %v", response)
		transaction := entity.Transaction{
			Id:          payment.OrderID,
			UserId:      UserId,
			Total:       *payment.Total,
			Status:      entity.Waiting,
			CreatedAt:   time.Now(),
			PaymentType: "payment-link",
			PaymentUrl:  response.PaymentUrl,
			ItemDetail:  payment.ItemDetail,
		}

		err := p.Repo.CreateTransaction(ctx, &transaction)
		if err != nil {
			return entity.PaymentLinkResponse{}, err
		}
	}

	return response, nil
}

func (p *PaymentService) CreateSnapPaymentMidtrans(ctx context.Context, payment *entity.PaymentDetail) (entity.PaymentSnapResponse, error) {
	// UserId := ctx.Value("user_id").(int)
	UserId := 1
	payment.OrderID = uuid.New()

	paymentRequest, err := PayloadRequestMidtransBuilder("snap", payment)
	if err != nil {
		return entity.PaymentSnapResponse{}, err
	}

	response, err := p.Repo.CreateSnapPaymentMidtrans(ctx, &paymentRequest)
	if err != nil {
		return entity.PaymentSnapResponse{}, err
	}

	// insert payment data if create midtrans transaction success
	if response.Status == http.StatusOK || response.Status == http.StatusCreated {
		log.Printf("New Snap Payment Created : %v", response)
		transaction := entity.Transaction{
			Id:          payment.OrderID,
			UserId:      UserId,
			Total:       *payment.Total,
			Status:      entity.Waiting,
			CreatedAt:   time.Now(),
			PaymentType: "snap-payment",
			PaymentUrl:  response.PaymentUrl,
			ItemDetail:  payment.ItemDetail,
		}

		err := p.Repo.CreateTransaction(ctx, &transaction)
		if err != nil {
			return entity.PaymentSnapResponse{}, err
		}
	}

	return response, nil
}

func (p *PaymentService) CreateCorePaymentMidtrans(ctx context.Context, payment *entity.PaymentDetail) (entity.PaymentCoreResponse, error) {
	// UserId := ctx.Value("user_id").(int)
	UserId := 1
	payment.OrderID = uuid.New()

	paymentRequest, err := PayloadRequestMidtransBuilder("core", payment)
	if err != nil {
		return entity.PaymentCoreResponse{}, err
	}

	response, err := p.Repo.CreateCorePaymentMidtrans(ctx, &paymentRequest)
	if err != nil {
		return entity.PaymentCoreResponse{}, err
	}

	// insert payment data if create midtrans transaction success
	if response.Status == http.StatusOK || response.Status == http.StatusCreated {
		log.Printf("New Core Payment Created : %v", response)
		transaction := entity.Transaction{
			Id:          payment.OrderID,
			UserId:      UserId,
			Total:       *payment.Total,
			Status:      entity.Waiting,
			CreatedAt:   time.Now(),
			PaymentType: string(payment.PaymentType),
			PaymentId:   response.PaymentId,
			ItemDetail:  payment.ItemDetail,
			// PaymentUrl:  response.VANumber,
		}

		err := p.Repo.CreateTransaction(ctx, &transaction)
		if err != nil {
			return entity.PaymentCoreResponse{}, err
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
