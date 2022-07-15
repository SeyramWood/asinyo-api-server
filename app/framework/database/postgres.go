package database

import (
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	_ "github.com/lib/pq"
	"log"
)

func postgresConnector(dBDriver string) *ent.Client {
	conf := config.DB()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Name,
		conf.Password,
	)
	client, err := ent.Open(dBDriver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	return client
}
