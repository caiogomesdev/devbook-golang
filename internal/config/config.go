package config

import (
	"os"

	"github.com/joho/godotenv"
)

type db struct {
  HOST string
  PORT string
  USER string
  PASSWORD string
  DATABASE string
}

type api struct {
  PORT string
}

type env struct{
  Db db
  Api api
}

var Env env

func loadEnvs(){
  Env = env{
    Db: db{
      HOST: os.Getenv("DB_HOST"),
      USER: os.Getenv("DB_USER"),
      PORT: os.Getenv("DB_PORT"),
      PASSWORD: os.Getenv("DB_PASS"),
      DATABASE: os.Getenv("DB_DATABASE"),
    },
    Api: api{
      PORT: os.Getenv("API_PORT"),
    },
  }
}

func Init(){
  godotenv.Load();
  loadEnvs()
  configureDb();
}
