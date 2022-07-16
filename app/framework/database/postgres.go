package database

import (
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/pkg/env"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func postgresConnector(dBDriver string) *ent.Client {
	conf := config.DB()
	var DbSSLMode string
	if os.Getenv("APP_ENV") == "production" {
		DbSSLMode = os.Getenv("DB_SSLMODE")
	} else {
		DbSSLMode = env.Get("DB_SSLMODE", "disable")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Name,
		conf.Password,
		DbSSLMode,
	)
	client, err := ent.Open(dBDriver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	return client
}
