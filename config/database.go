package config

import (
	"context"
	"crypto/tls"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(conf *DBConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conf.DatabaseURI).SetServerAPIOptions(serverAPI)
	db, err := mongo.Connect(context.TODO(), opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		return nil, err
	}

	return db, nil
}
