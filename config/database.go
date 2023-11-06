package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(conf *DBConfig) (*mongo.Client, error) {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.DatabaseURI))
	if err != nil {
		return nil, err
	}

	return db, nil
}
