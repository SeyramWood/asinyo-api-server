package config

import "github.com/SeyramWood/pkg/env"

type server struct {
	Prefork           bool
	CaseSensitive     bool
	StrictRouting     bool
	StreamRequestBody bool
	EnablePrintRoutes bool
	Concurrency       int64
	ServerHeader      string
	AppName           string
}

func Server() *server {
	return &server{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		StreamRequestBody: true,
		EnablePrintRoutes: true,
		Concurrency:       256 * 2048,
		ServerHeader:      "Asinyo Corporations",
		AppName:           env.Get("APP_NAME", "Asinyo API Server"),
	}
}
