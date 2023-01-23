package payment

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

func NewPaymentService(repo gateways.PaymentRepo, paymentGateway ...string) gateways.PaymentService {
	if paymentGateway == nil {
		return getServiceType(repo, config.Payment().Gateway)
	}
	return getServiceType(repo, paymentGateway[0])
}

func getServiceType(repo gateways.PaymentRepo, gateway string) gateways.PaymentService {
	switch gateway {
	case "paystack":
		return newPaystackService(repo)
	case "pay_on_delivery":
		return newPayOnDeliveryService(repo)
	default:
		return nil
	}
}
