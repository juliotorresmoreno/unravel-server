package crud

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/graphql-go/graphql"
	reflections "gopkg.in/oleiade/reflections.v1"
)

func getData(p graphql.ResolveParams, data Table) {
	campos, _ := reflections.Fields(data)
	alias, _ := reflections.Tags(data, "json")
	element := reflect.ValueOf(data).Elem()
	for _, campo := range campos {
		if ok, _ := reflections.HasField(data, campo); ok {
			value := p.Args[alias[campo]]
			if value != nil {
				switch element.FieldByName(campo).Type().String() {
				case "string":
					reflections.SetField(data, campo, value.(string))
				case "int":
					reflections.SetField(data, campo, value.(int))
				case "time.Time":
					if campo != "CreateAt" && campo != "UpdateAt" {
						t := strings.Split(value.(string), "-")
						if len(t) == 3 {
							y, _ := strconv.Atoi(t[0])
							m, _ := strconv.Atoi(t[1])
							d, _ := strconv.Atoi(t[2])
							f := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
							reflections.SetField(data, campo, f)
						}
					}
				}
			}
		}
	}
}

func validateUnique(orm *xorm.Engine, data interface{}, id int) error {
	campos, _ := reflections.Fields(data)
	valid, _ := reflections.Tags(data, "valid")
	for _, campo := range campos {
		if ok, _ := reflections.HasField(data, campo); ok {
			rules := strings.Split(valid[campo], " ")
			value, _ := reflections.GetField(data, campo)
			for i := range rules {
				if rules[i] == "unique" {
					cond := fmt.Sprintf("id != ? and state = 1 and %v = ?", campo)
					conteo, err := orm.Where(cond, id, fmt.Sprintf("%v", value)).Count(data)
					if err != nil {
						return err
					}
					if conteo > 0 {
						return fmt.Errorf("Ya existe el %v %v", campo, value)
					}
					break
				}
			}
		}
	}
	return nil
}
