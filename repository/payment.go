package repository

import (
	"context"

	"github.com/google/uuid"
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

func (p *PaymentRepository) CreateTransaction(ctx context.Context, payment *entity.PaymentDetail) error {
	_, err := p.DB.Database(DBName).Collection("transaction").InsertOne(ctx, bson.D{
		{Key: "_id", Value: payment.OrderID},
		{Key: "date", Value: payment.Date},
		{Key: "total", Value: payment.Total},
		{Key: "status", Value: payment.Status},
		{Key: "payment_type", Value: payment.PaymentType},
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
