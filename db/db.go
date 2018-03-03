package db

import (
	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/unravel-server/config"
	redis "gopkg.in/redis.v5"
)

var cache *redis.Client

// GetCache Obtiene la cache
func GetCache() *redis.Client {
	return cache
}

// GetXORM Obtiene el orm con acceso a la base de datos
func GetXORM() *xorm.Engine {
	var dsn string
	var err error
	var charset = "?charset=utf8&parseTime=true"
	var host = config.DB_HOST + ":" + config.DB_PORT
	if config.DB_PSWD != "" {
		dsn = config.DB_USER + ":" + config.DB_PSWD + "@tcp(" + host + ")/" + config.DB_DB + charset
	} else {
		dsn = config.DB_USER + "@tcp(" + host + ")/" + config.DB_DB + charset
	}
	orm, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return orm
}

func init() {
	cache = redis.NewClient(&redis.Options{
		Addr: config.REDIS_HOST + ":" + config.REDIS_PORT,
	})
}
