package config

import "github.com/SeyramWood/pkg/env"

// AppConfig - our app specific config
type app struct {
	Name      string
	Version   string
	AppURL    string
	ServerURL string
	Key       string
	TokenName string
}

type paystack struct {
	URL    string
	PubKey string
	SecKey string
	Email  string
	Domain string
}

func App() *app {
	if env.Get("APP_ENV", "local") == "local" {
		return &app{
			Name:      env.Get("APP_NAME", "My First API"),
			Version:   env.Get("APP_VERSION", "0.0.1"),
			AppURL:    env.Get("APP_URL", "http://127.0.0.1:8000"),
			ServerURL: env.Get("SERVER_URL", "127.0.0.1:8000"),
			Key:       env.Get("APP_KEY", "secretKEY5465"),
			TokenName: env.Get("API_TOKEN_NAME", "asinyo_remember"),
		}
	}
	return &app{
		Name:      env.GetProd("APP_NAME", "My First API"),
		Version:   env.GetProd("APP_VERSION", "0.0.1"),
		AppURL:    env.GetProd("APP_URL", "http://127.0.0.1:8000"),
		ServerURL: env.GetProd("SERVER_URL", "127.0.0.1:8000"),
		Key:       env.GetProd("APP_KEY", "secretKEY5465"),
		TokenName: env.GetProd("API_TOKEN_NAME", "asinyo_remember"),
	}
}

func Paystack() *paystack {
	if env.Get("APP_ENV", "local") == "local" {
		return &paystack{
			URL:    env.Get("PAYSTACK_URL", ""),
			PubKey: env.Get("PAYSTACK_PUB_KEY", ""),
			SecKey: env.Get("PAYSTACK_SEC_KEY", ""),
			Email:  env.Get("PAYSTACK_EMAIL", ""),
			Domain: env.Get("PAYSTACK_DOMAIN", ""),
		}
	}
	return &paystack{
		URL:    env.GetProd("PAYSTACK_URL", ""),
		PubKey: env.GetProd("PAYSTACK_PUB_KEY", ""),
		SecKey: env.GetProd("PAYSTACK_SEC_KEY", ""),
		Email:  env.GetProd("PAYSTACK_EMAIL", ""),
		Domain: env.GetProd("PAYSTACK_DOMAIN", ""),
	}
}
