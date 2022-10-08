package payment

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

func NewPaymentService(repo gateways.PaymentRepo) gateways.PaymentService {
	switch config.Payment().Gateway {
	case "paystack":
		return newPaystackService(repo)
	default:
		return nil
	}
}
