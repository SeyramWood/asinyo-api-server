package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Env map[string]string
var EnvProduction map[string]string

func Get(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func GetProduction(key, def string) string {
	if val, ok := EnvProduction[key]; ok {
		return val
	}
	return def
}

func Setup() {
	if os.Getenv("APP_ENV") != "production" {
		envFile := ".env"
		envProduction := ".env.production"
		var err error
		Env, err = godotenv.Read(envFile)
		if err != nil {
			log.Fatalln(err)
		}
		EnvProduction, err = godotenv.Read(envProduction)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
