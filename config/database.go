package config

import (
	"github.com/SeyramWood/pkg/env"
	"os"
)

// AppConfig - our app specific config
type db struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func DB() *db {
	var driver string
	if os.Getenv("APP_ENV") == "production" {
		driver = os.Getenv("DB_DRIVER")
	} else {
		driver = env.Get("DB_DRIVER", "sqlite")
	}
	switch driver {
	case "mysql":
		if os.Getenv("APP_ENV") == "production" {
			return &db{
				Driver:   driver,
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
			}
		}
		return &db{
			Driver:   driver,
			Host:     env.Get("DB_HOST", "127.0.0.1"),
			Port:     env.Get("DB_PORT", "3306"),
			Name:     env.Get("DB_DATABASE", "test_db"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
		}
	case "postgres":
		if os.Getenv("APP_ENV") == "production" {
			return &db{
				Driver:   driver,
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
			}

		}
		return &db{
			Driver:   driver,
			Host:     env.GetProduction("DB_HOST", "127.0.0.1"),
			Port:     env.GetProduction("DB_PORT", "5432"),
			Name:     env.GetProduction("DB_DATABASE", "test_db"),
			Username: env.GetProduction("DB_USERNAME", "postgres"),
			Password: env.GetProduction("DB_PASSWORD", ""),
		}
	}
	return nil
}
