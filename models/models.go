package models

import (
	_ "github.com/go-sql-driver/mysql"
	"../config"
	"github.com/asaskevich/govalidator"
	"regexp"
	"../lang/es"
	"strings"
	"github.com/go-xorm/xorm"
	"fmt"
	"gopkg.in/redis.v5"
)

var orm *xorm.Engine
var cache *redis.Client

var rDuplicateEntry *regexp.Regexp

type werror struct {
	Msg string
}

func (e werror) Error() string {
	return e.Msg
}

func init() {
	var dsn string
	if config.DB_PSWD != "" {
		dsn = config.DB_USER + ":" + config.DB_PSWD +
			"@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_DB +
			"?charset=utf8"
	} else {
		dsn = config.DB_USER +
			"@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_DB +
			"?charset=utf8&parseTime=true"
	}
	orm, _ = xorm.NewEngine("mysql", dsn)

	govalidator.TagMap["alphaSpaces"] = govalidator.Validator(func(str string) bool {
		patterm, _ := regexp.Compile("^([a-zA-Z]+( ){0,1}){1,}$")
		return patterm.MatchString(str)
	})
	govalidator.TagMap["password"] = govalidator.Validator(func(str string) bool {
		return len(str) > 4
	})
	rDuplicateEntry, _ = regexp.Compile("Error 1062")
	cache = redis.NewClient(&redis.Options{
		Addr: config.REDIS_HOST + ":" + config.REDIS_PORT,
	})
}

func GetCache() *redis.Client {
	return cache
}

func GetXORM() *xorm.Engine {
	return orm
}

func normalize(Error error, data interface{}) error {
	var message string
	if rDuplicateEntry.MatchString(Error.Error()) {
		fmt.Println(Error)
		values := strings.Split(Error.Error(), "'")
		message = strings.Replace(es.DuplicateEntry, "{campo}", strings.Split(values[3], "_")[2], 1)
		message = strings.Replace(message, "{valor}", values[1], 1)
		return werror{Msg:message}
	}
	return Error
}