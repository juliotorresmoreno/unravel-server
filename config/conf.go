package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

var (
	DB_HOST          string = "localhost"
	DB_USER          string
	DB_PORT          string
	DB_PSWD          string
	DB_DB            string
	REDIS_HOST       string = "localhost"
	REDIS_PORT       string = "6379"
	PATH             string
	SESSION_DURATION int = 7200
	PORT             int = 80
	READ_TIMEOUT     time.Duration
	USERNAME         string
	PASSWORD         string
	SERVIDOR         string
	PUERTO           int
	DATABASE         string
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
	SessionDuration string
	Port            string
	ReadTimeout     string
	Username        string
	Password        string
	Servidor        string
	Puerto          string
	Database        string
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
	SESSION_DURATION, _ = strconv.Atoi(data.SessionDuration)
	PORT, _ = strconv.Atoi(data.Port)
	READ_TIMEOUT, _ = time.ParseDuration(data.ReadTimeout)
	USERNAME = data.Username
	PASSWORD = data.Password
	SERVIDOR = data.Servidor
	PUERTO, _ = strconv.Atoi(data.Puerto)
	DATABASE = data.Database
}
