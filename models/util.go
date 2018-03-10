package models

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/lang/es"
	redis "gopkg.in/redis.v5"
)

var rDuplicateEntry *regexp.Regexp
var cache *redis.Client

func init() {
	govalidator.TagMap["alphaSpaces"] = govalidator.Validator(func(str string) bool {
		patterm, _ := regexp.Compile("^([a-zA-Z]+( ){0,1}){1,}$")
		return patterm.MatchString(str)
	})
	govalidator.TagMap["username"] = govalidator.Validator(func(str string) bool {
		patterm, _ := regexp.Compile("^[a-zA-Z0-9_]{3,}$")
		return patterm.MatchString(str)
	})
	govalidator.TagMap["password"] = govalidator.Validator(func(str string) bool {
		return len(str) > 4
	})
	govalidator.TagMap["precio_hora"] = govalidator.Validator(func(str string) bool {
		t, _ := strconv.Atoi(str)
		return t >= 5
	})

	rDuplicateEntry, _ = regexp.Compile("Error 1062")
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
	var orm = db.GetXORM()
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
	var orm = db.GetXORM()
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
