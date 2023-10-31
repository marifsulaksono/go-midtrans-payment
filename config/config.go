package config

import "os"

const (
	MidtransServerKey   = "MIDTRANS_SERVER_KEY"
	MidtransSandboxLink = "MIDTRANS_SANDBOX_LINK"
)

type PaymentConfig struct {
	ServerKey   string
	SandboxLink string
}

func GetPaymentConfig() *PaymentConfig {
	return &PaymentConfig{
		ServerKey:   os.Getenv(MidtransServerKey),
		SandboxLink: os.Getenv(MidtransSandboxLink),
	}
}
