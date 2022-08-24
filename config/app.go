package config

import (
	"os"

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
