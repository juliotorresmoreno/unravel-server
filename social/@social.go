package social

import (
	"../config"
	"gopkg.in/mgo.v2"
)

var username = config.MONGO_USER
var password = config.MONGO_PSWD
var servidor = config.MONGO_HOST
var puerto = config.MONGO_PORT
var database = config.MONGO_DB
var url = "mongodb://" + username + ":" + password + "@" + servidor + ":" + puerto + "/" + database

func Add(collection string, data interface{}) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	noticias := session.DB(database).C(collection)
	err = noticias.Insert(data)
	return err
}
