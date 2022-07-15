package database

import (
	"database/sql"
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func mysqlConnector() *ent.Client {
	conf := config.DB("mysql")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	//time.Sleep(time.Second * 10)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	return client
}
func MysqlConnector2() *sql.DB {
	conf := config.DB("mysql")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()

	return db

}
