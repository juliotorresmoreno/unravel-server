package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/juliotorresmoreno/unravel-server/db"
)

//Friend Entidad que representa una amistad
type Friend struct {
	Usuario    string    `json:"usuario"`
	Nombres    string    `json:"nombres"`
	Apellidos  string    `json:"apellidos"`
	FullName   string    `json:"fullname"`
	Estado     string    `json:"estado"`
	Registrado time.Time `json:"registrado"`
	Conectado  bool      `json:"conectado"`
	Relacion   *Relacion `json:"relacion"`
}

//IsFriend Determina su dos usuarios son amigos
func IsFriend(usuario string, amigo string) int8 {
	var relaciones = make([]Relacion, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var str = "(usuario_solicita = ? and usuario_solicitado = ?) or (usuario_solicita = ? and usuario_solicitado = ?)"
	orm.Where(str, usuario, amigo, amigo, usuario).Find(&relaciones)
	if len(relaciones) == 1 {
		return int8(relaciones[0].EstadoRelacion)
	}
	return EstadoDesconocido
}

//GetFriends Obtiene el listado de amigos
func GetFriends(usuario string) ([]*Friend, error) {
	var defecto = make([]*Friend, 0)
	var relaciones = make([]Relacion, 0)
	var users = make([]User, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var str = "usuario_solicita = ? or usuario_solicitado = ?"
	if err := orm.Where(str, usuario, usuario).Find(&relaciones); err != nil {
		return defecto, errors.New("Error desconocido")
	}
	if len(relaciones) > 0 {
		var data = ""
		for _, el := range relaciones {
			if el.UsuarioSolicitado == usuario {
				data += "\"" + el.UsuarioSolicita + "\", "
			} else {
				data += "\"" + el.UsuarioSolicitado + "\", "
			}
		}
		data = data[0 : len(data)-2]
		str = "Usuario in (" + data + ")"
		if err := orm.Where(str).Find(&users); err != nil {
			return defecto, errors.New("Error desconocido")
		}
	}
	return listUserToListFriends(users, relaciones), nil
}

func listUserToListFriends(users []User, relacion []Relacion) []*Friend {
	var lengthUsers = len(users)
	var lengthRelacion = len(relacion)
	var list = make([]*Friend, lengthUsers)

	for i := 0; i < lengthUsers; i++ {
		list[i] = &Friend{
			Usuario:    users[i].Usuario,
			Nombres:    users[i].Nombres,
			Apellidos:  users[i].Apellidos,
			FullName:   users[i].FullName,
			Estado:     "Desconocido",
			Registrado: users[i].CreateAt,
		}
		for j := 0; j < lengthRelacion; j++ {
			solicita := relacion[j].UsuarioSolicita
			solicitado := relacion[j].UsuarioSolicitado
			if users[i].Usuario == solicita || users[i].Usuario == solicitado {
				list[i].Relacion = &relacion[j]
				if int8(relacion[j].EstadoRelacion) == EstadoSolicitado {
					list[i].Estado = "Solicitado"
				} else {
					list[i].Estado = "Amigos"
				}
			}
		}
	}
	return list
}

//FindUsers Busca los usuarios
func FindUsers(usuarios []string) (*[]User, error) {
	var users = make([]User, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var str string
	str = "Usuario in ('" + strings.Join(usuarios, "', '") + "')"
	if err := orm.Where(str).Find(&users); err != nil {
		return &users, err
	}
	return &users, nil
}

//FindUser Busca un unico usuario
func FindUser(session string, query string, usuario string) ([]*Friend, error) {
	var users = make([]User, 0)
	var relaciones = make([]Relacion, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var str string
	if query != "" {
		w := strings.Split(query, " ")
		str = "usuario != ? AND (false"
		for _, v := range w {
			str = str + " OR full_name LIKE '%" + v + "%'"
		}
		str = str + ")"
	} else if usuario != "" {
		str = "Usuario != ? AND Usuario = '" + usuario + "'"
	} else {
		str = "Usuario != ?"
	}

	if err := orm.Where(str, session).Find(&users); err != nil {
		return make([]*Friend, 0), err
	}

	str = "usuario_solicita = ? OR usuario_solicitado = ?"
	if err := orm.Where(str, session, session).Find(&relaciones); err != nil {
		return make([]*Friend, 0), err
	}

	return listUserToListFriends(users, relaciones), nil
}

//RejectFriends rechazar amistad
func RejectFriends(session string, usuario string) (int64, error) {
	if session == "" || usuario == "" {
		return 0, nil
	}
	var orm = db.GetXORM()
	defer orm.Close()
	var relacion Relacion
	var aff int64
	var str = "(usuario_solicita = \"%s\" AND usuario_solicitado = \"%s\") OR (usuario_solicita = \"%s\" AND usuario_solicitado = \"%s\")"
	result, err := orm.Exec(fmt.Sprintf("DELETE FROM "+relacion.TableName()+" WHERE "+str, session, usuario, usuario, session))
	if result != nil {
		aff, _ = result.RowsAffected()
	}
	return aff, err
}
