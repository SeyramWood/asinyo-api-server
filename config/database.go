package config

import (
	"github.com/SeyramWood/pkg/env"
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
	driver := env.Get("DB_DRIVER", "sqlite")
	switch driver {
	case "mysql":
		if env.Get("APP_ENV", "local") == "local" {
			return &db{
				Driver:   driver,
				Host:     env.Get("DB_HOST", "127.0.0.1"),
				Port:     env.Get("DB_PORT", "3306"),
				Name:     env.Get("DB_DATABASE", "test_db"),
				Username: env.Get("DB_USERNAME", "root"),
				Password: env.Get("DB_PASSWORD", ""),
			}
		}
		return &db{
			Driver:   driver,
			Host:     env.GetProduction("DB_HOST", "127.0.0.1"),
			Port:     env.GetProduction("DB_PORT", "3306"),
			Name:     env.GetProduction("DB_DATABASE", "test_db"),
			Username: env.GetProduction("DB_USERNAME", "root"),
			Password: env.GetProduction("DB_PASSWORD", ""),
		}
	case "postgres":
		if env.Get("APP_ENV", "local") == "local" {
			return &db{
				Driver:   driver,
				Host:     env.Get("DB_HOST", "127.0.0.1"),
				Port:     env.Get("DB_PORT", "5432"),
				Name:     env.Get("DB_DATABASE", "test_db"),
				Username: env.Get("DB_USERNAME", "postgres"),
				Password: env.Get("DB_PASSWORD", ""),
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
