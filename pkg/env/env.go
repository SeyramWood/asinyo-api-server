package env

import "github.com/joho/godotenv"

var Env map[string]string
var EnvProd map[string]string

func Get(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func GetProd(key, def string) string {
	if val, ok := EnvProd[key]; ok {
		return val
	}
	return def
}

func Setup() {
	envFile := ".env"
	envFileProd := ".env.production"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}
	EnvProd, err = godotenv.Read(envFileProd)
	if err != nil {
		panic(err)
	}

}
