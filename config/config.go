package config

import "os"

const (
	MidtransServerKey       = "MIDTRANS_SERVER_KEY"
	MidtransSandboxLink     = "MIDTRANS_SANDBOX_LINK"
	MidtransSnapSandboxLink = "MIDTRANS_SNAP_SANDBOX_LINK"
)

type PaymentConfig struct {
	ServerKey       string
	SandboxLink     string
	SnapSandboxLink string
}

func GetPaymentConfig() *PaymentConfig {
	return &PaymentConfig{
		ServerKey:       os.Getenv(MidtransServerKey),
		SandboxLink:     os.Getenv(MidtransSandboxLink),
		SnapSandboxLink: os.Getenv(MidtransSnapSandboxLink),
	}
}
