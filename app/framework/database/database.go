package database

import (
	"github.com/SeyramWood/ent"
)

type Adapter struct {
	DB *ent.Client
}

func NewDB(driver string) *Adapter {
	switch driver {
	case "mysql":
		return &Adapter{DB: mysqlConnector()}
	}
	return nil
}
