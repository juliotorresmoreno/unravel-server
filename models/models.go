package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"../config"
	"github.com/asaskevich/govalidator"
	"regexp"
	"../lang/es"
	"strings"
)

//var engine *xorm.Engine
var engine *gorm.DB
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
			"?charset=utf8"
	}
	engine, _ = gorm.Open("mysql", dsn)
	govalidator.TagMap["alphaSpaces"] = govalidator.Validator(func(str string) bool {
		patterm, _ := regexp.Compile("^([a-zA-Z]+( ){0,1}){1,}$")
		return patterm.MatchString(str)
	})
	govalidator.TagMap["password"] = govalidator.Validator(func(str string) bool {
		return len(str) > 4
	})
	rDuplicateEntry, _ = regexp.Compile("Error 1062")
}

func normalize(Error error, data interface{}) error {
	bError := []byte(Error.Error())
	if rDuplicateEntry.Match(bError) {
		luser, _ := regexp.Compile("'[a-zA-Z]*'$")
		c:=luser.Find(bError)
		campo := strings.Replace(string(c), "'", "", -1)
		m:= strings.Replace(es.DuplicateEntry, "{campo}", campo, 1)
		return werror{Msg:m}
	}
	return Error
}