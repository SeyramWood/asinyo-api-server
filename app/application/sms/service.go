package sms

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

func NewSMSService() gateways.SMSService {
	switch config.SMS().Gateway {
	case "arkesel":
		return newArkeselService()
	default:
		return nil
	}
}
