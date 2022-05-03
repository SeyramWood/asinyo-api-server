package database

import (
	"context"
	"fmt"
	"log"

	"database/sql"

	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
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
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	// defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
