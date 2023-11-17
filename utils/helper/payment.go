package helper

import "github.com/marifsulaksono/go-midtrans-payment/entity"

// checking valid payment type
func IsValidPaymentType(paymentType entity.Type) bool {
	validTypes := map[entity.Type]bool{
		entity.Bank:    true,
		entity.Mandiri: true,
		entity.Permata: true,
		entity.Store:   true,
		entity.Alku:    true,
		entity.Kred:    true,
		entity.Gopay:   true,
		entity.Qris:    true,
	}

	return validTypes[paymentType]
}
