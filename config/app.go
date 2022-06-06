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

func App() *app {
	return &app{
		Name:      env.Get("APP_NAME", "My First API"),
		Version:   env.Get("APP_VERSION", "0.0.1"),
		AppURL:    env.Get("APP_URL", "http://127.0.0.1:8000"),
		ServerURL: env.Get("SERVER_URL", "127.0.0.1:8000"),
		Key:       env.Get("APP_KEY", "secretKEY5465"),
		TokenName: env.Get("API_TOKEN_NAME", "asinyo_remember"),
	}
}
