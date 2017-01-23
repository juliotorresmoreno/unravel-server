package models

import (
	"errors"
	"regexp"
	"strings"

	"../config"
	"../lang/es"
	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/redis.v5"
)

var orm *xorm.Engine
var cache *redis.Client

var rDuplicateEntry *regexp.Regexp

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

// GetCache Obtiene la cache
func GetCache() *redis.Client {
	return cache
}

// GetXORM Obtiene el orm con acceso a la base de datos
func GetXORM() *xorm.Engine {
	return orm
}

func normalize(Error error, data interface{}) error {
	var message string
	if rDuplicateEntry.MatchString(Error.Error()) {
		var values = strings.Split(Error.Error(), "'")
		var campo = strings.Split(values[3], "_")[2]
		message = strings.Replace(es.DuplicateEntry, "{campo}", campo, 1)
		message = strings.Replace(message, "{valor}", values[1], 1)
		return errors.New(campo + ": " + message)
	}
	return Error
}

// Update valida y actualiza un nuevo registro en base de datos
func Update(id uint, self interface{}) (int64, error) {
	_, err := govalidator.ValidateStruct(self)
	if err != nil {
		return 0, normalize(err, self)
	}

	affected, err := orm.Id(id).Update(self)
	if err != nil {
		return affected, normalize(err, self)
	}
	return affected, nil
}

// Add valida y crea un nuevo registro en base de datos
func Add(self interface{}) (int64, error) {
	_, err := govalidator.ValidateStruct(self)
	if err != nil {
		return 0, normalize(err, self)
	}
	affected, err := orm.Insert(self)
	if err != nil {
		return affected, normalize(err, self)
	}
	return affected, nil
}

type model interface {
	TableName() string
	getID() uint
}
