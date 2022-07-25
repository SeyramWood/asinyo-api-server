package payment

import "github.com/SeyramWood/app/adapters/gateways"

func NewPaymentService(repo gateways.PaymentRepo, gateway string) gateways.PaymentService {
	switch gateway {
	case "paystack":
		return NewPaystackService(repo)
	default:
		return nil
	}
}
