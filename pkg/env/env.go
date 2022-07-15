package env

import (
	"github.com/joho/godotenv"
	"log"
)

var Env map[string]string
var EnvDocker map[string]string

func Get(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func GetDocker(key, def string) string {
	if val, ok := EnvDocker[key]; ok {
		return val
	}
	return def
}

func Setup() {
	envFile := ".env"
	envFileDocker := ".env.docker"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		log.Fatalln(err)
	}
	EnvDocker, err = godotenv.Read(envFileDocker)
	if err != nil {
		log.Fatalln(err)
	}

}
