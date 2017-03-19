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
	HOSTNAME         string
	PORT             int = 80
	PORT_SSL         int = 443
	CERT_FILE        string
	KEY_FILE         string
	READ_TIMEOUT     time.Duration
	MONGO_HOST       string
	MONGO_USER       string
	MONGO_PORT       string
	MONGO_PSWD       string
	MONGO_DB         string
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
	Hostname        string
	Port            string
	PortSsl         string
	CertFile        string
	KeyFile         string
	ReadTimeout     string
	MongoHost       string
	MongoUser       string
	MongoPort       string
	MongoPswd       string
	MongoDb         string
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

	HOSTNAME = data.Hostname
	PORT, _ = strconv.Atoi(data.Port)
	PORT_SSL, _ = strconv.Atoi(data.PortSsl)
	CERT_FILE = data.CertFile
	KEY_FILE = data.KeyFile

	READ_TIMEOUT, _ = time.ParseDuration(data.ReadTimeout)

	MONGO_HOST = data.MongoHost
	MONGO_USER = data.MongoUser
	MONGO_PORT = data.MongoPort
	MONGO_PSWD = data.MongoPswd
	MONGO_DB = data.MongoDb
}
