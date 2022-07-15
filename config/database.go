package config

import "github.com/SeyramWood/pkg/env"

// AppConfig - our app specific config
type db struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func DB(driver string) *db {
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
			Host:     env.GetDocker("DB_HOST", "localhost"),
			Port:     env.GetDocker("DB_PORT", "3306"),
			Name:     env.GetDocker("DB_DATABASE", "test_db"),
			Username: env.GetDocker("DB_USERNAME", "root"),
			Password: env.GetDocker("DB_PASSWORD", ""),
		}
	}
	return nil
}
