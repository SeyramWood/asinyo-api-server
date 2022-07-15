package database

import (
	"database/sql"
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/pkg/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

type Adapter struct {
	DB *ent.Client
}

func NewDB() *Adapter {
	dBDriver := env.Get("DB_DRIVER", "sqlite")
	switch dBDriver {
	case "mysql":
		return &Adapter{DB: mysqlConnector(dBDriver)}
	case "postgres":
		return &Adapter{DB: postgresConnector(dBDriver)}
	}
	return nil
}

func Connect() *sql.DB {
	conf := config.DB()
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Name,
			conf.Password,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Fatalf(err.Error())
		}
		return db
	default:
		mysqlDSN := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		)
		db, err := sql.Open(conf.Driver, mysqlDSN)
		if err != nil {
			log.Fatalf(err.Error())
		}
		return db
	}

}
