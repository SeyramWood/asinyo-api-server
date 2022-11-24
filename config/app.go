package config

import (
	"os"
	"sync"

	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/pkg/env"
)

// AppConfig - our app specific config
type app struct {
	Name      string
	Version   string
	AppURL    string
	Key       string
	TokenName string
	PORT      string
}
type payment struct {
	Gateway string
}
type sms struct {
	Sender  string
	Gateway string
}

type logistic struct {
	Gateway   string
	WG        *sync.WaitGroup
	DataChan  chan *ent.Order
	DoneChan  chan bool
	ErrorChan chan error
}

type mailer struct {
	Mailer      string
	Host        string
	Port        string
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type paystack struct {
	URL    string
	PubKey string
	SecKey string
	Email  string
	Domain string
}
type uploadcare struct {
	PubKey string
	SecKey string
	URL    string
}

type arkesel struct {
	APIKey string
	URL    string
}
type tookan struct {
	APIKey string
	URL    string
}

type google struct {
	APIKey string
}

func App() *app {
	if os.Getenv("APP_ENV") == "production" {
		return &app{
			Name:      os.Getenv("APP_NAME"),
			Version:   os.Getenv("APP_VERSION"),
			AppURL:    os.Getenv("APP_URL"),
			Key:       os.Getenv("APP_KEY"),
			TokenName: os.Getenv("API_TOKEN_NAME"),
			PORT:      os.Getenv("SERVER_PORT"),
		}
	}
	return &app{
		Name:      env.Get("APP_NAME", "My First API"),
		Version:   env.Get("APP_VERSION", "0.0.1"),
		AppURL:    env.Get("APP_URL", "http://127.0.0.1:9000"),
		Key:       env.Get("APP_KEY", "secretKEY5465"),
		TokenName: env.Get("API_TOKEN_NAME", "asinyo_remember"),
		PORT:      env.Get("SERVER_PORT", "9000"),
	}
}

func Logistic() *logistic {
	if os.Getenv("APP_ENV") == "production" {
		return &logistic{
			Gateway: os.Getenv("LOGISTIC_GATEWAY"),
		}
	}
	return &logistic{
		Gateway: env.Get("LOGISTIC_GATEWAY", "tookan"),
	}
}

func Payment() *payment {
	if os.Getenv("APP_ENV") == "production" {
		return &payment{
			Gateway: os.Getenv("PAYMENT_GATEWAY"),
		}
	}
	return &payment{
		Gateway: env.Get("PAYMENT_GATEWAY", "paystack"),
	}
}

func Mailer() *mailer {
	if os.Getenv("APP_ENV") == "production" {
		return &mailer{
			Mailer:      os.Getenv("MAIL_MAILER"),
			Host:        os.Getenv("MAIL_HOST"),
			Port:        os.Getenv("MAIL_PORT"),
			Username:    os.Getenv("MAIL_USERNAME"),
			Password:    os.Getenv("MAIL_PASSWORD"),
			Encryption:  os.Getenv("MAIL_ENCRYPTION"),
			FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
			FromName:    os.Getenv("MAIL_FROM_NAME"),
		}
	}
	return &mailer{
		Mailer:      env.Get("MAIL_MAILER", "smtp"),
		Host:        env.Get("MAIL_HOST", ""),
		Port:        env.Get("MAIL_PORT", ""),
		Username:    env.Get("MAIL_USERNAME", ""),
		Password:    env.Get("MAIL_PASSWORD", ""),
		Encryption:  env.Get("MAIL_ENCRYPTION", ""),
		FromAddress: env.Get("MAIL_FROM_ADDRESS", "info@asinyo.com"),
		FromName:    env.Get("MAIL_FROM_NAME", "Asinyo Agro-Cormmerce"),
	}
}

func SMS() *sms {
	if os.Getenv("APP_ENV") == "production" {
		return &sms{
			Sender:  os.Getenv("SMS_SENDER"),
			Gateway: os.Getenv("SMS_GATEWAY"),
		}
	}
	return &sms{
		Sender:  env.Get("SMS_SENDER", "Asinyo"),
		Gateway: env.Get("SMS_GATEWAY", "arkesel"),
	}
}

func Paystack() *paystack {
	if os.Getenv("APP_ENV") == "production" {
		return &paystack{
			URL:    os.Getenv("PAYSTACK_URL"),
			PubKey: os.Getenv("PAYSTACK_PUB_KEY"),
			SecKey: os.Getenv("PAYSTACK_SEC_KEY"),
			Email:  os.Getenv("PAYSTACK_EMAIL"),
			Domain: os.Getenv("PAYSTACK_DOMAIN"),
		}
	}
	return &paystack{
		URL:    env.Get("PAYSTACK_URL", ""),
		PubKey: env.Get("PAYSTACK_PUB_KEY", ""),
		SecKey: env.Get("PAYSTACK_SEC_KEY", ""),
		Email:  env.Get("PAYSTACK_EMAIL", ""),
		Domain: env.Get("PAYSTACK_DOMAIN", ""),
	}

}

func Uploadcare() *uploadcare {
	if os.Getenv("APP_ENV") == "production" {
		return &uploadcare{
			PubKey: os.Getenv("UPLOADCARE_PUB_KEY"),
			SecKey: os.Getenv("UPLOADCARE_SEC_KEY"),
			URL:    os.Getenv("UPLOADCARE_URL"),
		}
	}
	return &uploadcare{
		PubKey: env.Get("UPLOADCARE_PUB_KEY", ""),
		SecKey: env.Get("UPLOADCARE_SEC_KEY", ""),
		URL:    env.Get("UPLOADCARE_URL", ""),
	}

}

func Arkesel() *arkesel {
	if os.Getenv("APP_ENV") == "production" {
		return &arkesel{
			APIKey: os.Getenv("ARKESEL_API_KEY"),
			URL:    os.Getenv("ARKESEL_URL"),
		}
	}
	return &arkesel{
		APIKey: env.Get("ARKESEL_API_KEY", ""),
		URL:    env.Get("ARKESEL_URL", ""),
	}
}

func Tookan() *tookan {
	if os.Getenv("APP_ENV") == "production" {
		return &tookan{
			APIKey: os.Getenv("TOOKAN_API_KEY"),
			URL:    os.Getenv("TOOKAN_URL"),
		}
	}
	return &tookan{
		APIKey: env.Get("TOOKAN_API_KEY", ""),
		URL:    env.Get("TOOKAN_URL", ""),
	}
}

func Google() *google {
	if os.Getenv("APP_ENV") == "production" {
		return &google{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		}
	}
	return &google{
		APIKey: env.Get("GOOGLE_API_KEY", ""),
	}
}
