package config

import (
	"fmt"
	"os"
)

const (
	MidtransServerKey       = "MIDTRANS_SERVER_KEY"
	MidtransSandboxLink     = "MIDTRANS_SANDBOX_LINK"
	MidtransSnapSandboxLink = "MIDTRANS_SNAP_SANDBOX_LINK"

	MongoDBUsername = "MONGODB_USERNAME"
	MongoDBPassword = "MONGODB_PASSWORD"
	MongoDBCluster  = "MONGODB_CLUSTER"
	MongoDBName     = "MONGODB_NAME"
)

type PaymentConfig struct {
	ServerKey       string
	SandboxLink     string
	SnapSandboxLink string
}

type DBConfig struct {
	DatabaseURI string
}

func GetPaymentConfig() *PaymentConfig {
	return &PaymentConfig{
		ServerKey:       os.Getenv(MidtransServerKey),
		SandboxLink:     os.Getenv(MidtransSandboxLink),
		SnapSandboxLink: os.Getenv(MidtransSnapSandboxLink),
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		DatabaseURI: fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/?retryWrites=true&w=majority",
			os.Getenv(MongoDBUsername), os.Getenv(MongoDBPassword), os.Getenv(MongoDBCluster)),
	}
}
