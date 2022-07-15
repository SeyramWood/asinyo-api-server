package database

import (
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func mysqlConnector(dBDriver string) *ent.Client {
	conf := config.DB()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)
	client, err := ent.Open(dBDriver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	return client
}
