package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-midtrans-payment/config"
	"github.com/marifsulaksono/go-midtrans-payment/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBName string = "mongo_store"

type PaymentRepository struct {
	DB *mongo.Client
}

func NewPaymentRepository(db *mongo.Client) PaymentRepository {
	return PaymentRepository{DB: db}
}

func (p *PaymentRepository) CreateLinkPaymentMidtrans(ctx context.Context, paymentRequest *entity.MidtransPayloadRequest) (entity.PaymentLinkResponse, error) {
	conf := config.GetPaymentConfig()
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))
	response, err := RequestMidtransHitter(authString, conf.SandboxLink, paymentRequest)
	if err != nil {
		return entity.PaymentLinkResponse{}, err
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return entity.PaymentLinkResponse{}, err
	}

	var result entity.PaymentLinkResponse
	result.Status = response.StatusCode
	result.OrderId = fmt.Sprint(responseJSON["order_id"])
	result.PaymentUrl = fmt.Sprint(responseJSON["payment_url"])

	return result, nil
}

func (p *PaymentRepository) CreateSnapPaymentMidtrans(ctx context.Context, paymentRequest *entity.MidtransPayloadRequest) (entity.PaymentSnapResponse, error) {
	conf := config.GetPaymentConfig()
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))
	response, err := RequestMidtransHitter(authString, conf.SnapSandboxLink, paymentRequest)
	if err != nil {
		return entity.PaymentSnapResponse{}, err
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return entity.PaymentSnapResponse{}, err
	}

	var result entity.PaymentSnapResponse
	result.Status = response.StatusCode
	result.Token = fmt.Sprint(responseJSON["token"])
	result.PaymentUrl = fmt.Sprint(responseJSON["redirect_url"])

	return result, nil
}

func (p *PaymentRepository) CreateCorePaymentMidtrans(ctx context.Context, paymentRequest *entity.MidtransPayloadRequest) (entity.PaymentCoreResponse, error) {
	conf := config.GetPaymentConfig()
	authString := base64.StdEncoding.EncodeToString([]byte(conf.ServerKey + ":"))
	response, err := RequestMidtransHitter(authString, conf.CoreSandboxLink, paymentRequest)
	if err != nil {
		return entity.PaymentCoreResponse{}, err
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return entity.PaymentCoreResponse{}, err
	}

	var result entity.PaymentCoreResponse
	if result.Status, err = strconv.Atoi(fmt.Sprint(responseJSON["status_code"])); err != nil {
		return entity.PaymentCoreResponse{}, err
	}
	result.PaymentId = fmt.Sprint(responseJSON["transaction_id"])
	result.VANumber = responseJSON["va_numbers"]

	return result, nil
}

func (p *PaymentRepository) CreateTransaction(ctx context.Context, payment *entity.Transaction) error {
	_, err := p.DB.Database(DBName).Collection("transactions").InsertOne(ctx, bson.D{
		{Key: "_id", Value: payment.Id},
		{Key: "user_id", Value: payment.UserId},
		{Key: "total", Value: payment.Total},
		{Key: "status", Value: payment.Status},
		{Key: "created_at", Value: payment.CreatedAt},
		{Key: "payment_type", Value: payment.PaymentType},
		{Key: "payment_url", Value: payment.PaymentUrl},
		{Key: "payment_id", Value: payment.PaymentId},
		{Key: "payment_time", Value: payment.PaymentTime},
		{Key: "item_details", Value: payment.ItemDetail},
	})

	return err
}

func (p *PaymentRepository) UpdateTransaction(ctx context.Context, id string, status entity.Status) error {
	objId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: primitive.Binary{Subtype: 0, Data: objId[:]}}}
	statusUpdate := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}}

	var trx entity.PaymentDetail
	err = p.DB.Database(DBName).Collection("transaction").FindOne(ctx, filter).Decode(&trx)
	if err != nil {
		return err
	}

	_, err = p.DB.Database(DBName).Collection("transaction").UpdateOne(ctx, filter, statusUpdate)
	return err
}
