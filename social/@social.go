package social

import (
	"../config"
	"gopkg.in/mgo.v2"
)

var username = config.USERNAME
var password = config.PASSWORD
var servidor = config.SERVIDOR
var puerto = string(config.PUERTO)
var database = config.DATABASE
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
