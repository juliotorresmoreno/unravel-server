package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	DB_HOST string = "localhost"
	DB_USER string
	DB_PORT string
	DB_PSWD string
	DB_DB   string

	REDIS_HOST string = "localhost"
	REDIS_PORT string = "6379"

	PATH string

	SESSION_DURATION int = 7200
)

type configuration struct {
	DbHost          string
	DbUser          string
	DbPort          string
	DbPswd          string
	DbDb            string
	RedisHost       string
	RedisPort       string
	Path            string
	SessionDuration int
}

func init() {
	var text, err = ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	var data = &configuration{}
	err = json.Unmarshal(text, data)
	if err != nil {
		panic(err)
	}
	DB_HOST = data.DbHost
	DB_USER = data.DbUser
	DB_PORT = data.DbPort
	DB_PSWD = data.DbPswd
	DB_DB = data.DbDb
	REDIS_HOST = data.RedisHost
	REDIS_PORT = data.RedisPort
	PATH = data.Path
	SESSION_DURATION = data.SessionDuration
}
