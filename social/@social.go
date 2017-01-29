package social

import (
	"../config"
	"github.com/asaskevich/govalidator"
	"gopkg.in/mgo.v2"
)

var username = config.MONGO_USER
var password = config.MONGO_PSWD
var servidor = config.MONGO_HOST
var puerto = config.MONGO_PORT
var database = config.MONGO_DB
var url = "mongodb://" + username + ":" + password + "@" + servidor + ":" + puerto + "/" + database

// Add agrega un elemento a la coleccion
func Add(collection string, data interface{}) error {
	var _, err = govalidator.ValidateStruct(data)
	if err != nil {
		return err
	}
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

// Update actualiza un elemento a la coleccion
func Update(collection string, id interface{}, data interface{}) error {
	var _, err = govalidator.ValidateStruct(data)
	if err != nil {
		return err
	}
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	noticias := session.DB(database).C(collection)
	err = noticias.UpdateId(id, data)
	return err
}

// GetSocial obtiene la base de datos
func GetSocial() (*mgo.Session, *mgo.Database, error) {
	var session, err = mgo.Dial(url)
	if err != nil {
		return session, nil, err
	}
	return session, session.DB(database), nil
}
