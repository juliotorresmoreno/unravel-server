package crud

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"
	"github.com/graphql-go/graphql"
	"gopkg.in/oleiade/reflections.v1"
)

//GraphQLGet List
func GraphQLGet(p graphql.ResolveParams, orm *xorm.Engine, data interface{}) error {
	id := p.Args["id"].(int)
	if id != 0 {
		return orm.Where("id = ?", id).And("state = 1").Find(data)
	}
	return orm.Where("state = 1").Find(data)
}

//GraphQLPut Add
func GraphQLPut(p graphql.ResolveParams, orm *xorm.Engine, data Table) (bool, error) {
	getData(p, data)
	reflections.SetField(data, "State", 1)
	if _, err := govalidator.ValidateStruct(data); err != nil {
		return false, err
	}
	if err := validateUnique(orm, data, 0); err != nil {
		return false, err
	}
	if _, err := orm.Insert(data); err != nil {
		return false, err
	}
	return true, nil
}

//GraphQLPost Modificar
func GraphQLPost(p graphql.ResolveParams, orm *xorm.Engine, data Table) (bool, error) {
	id, _ := p.Args["id"].(int)
	orm.Id(id).Where("state = 1").Get(data)
	getData(p, data)
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		return false, err
	}
	if err := validateUnique(orm, data, id); err != nil {
		return false, err
	}
	_, err = orm.Id(id).Where("state = 1").Update(data)
	if err != nil {
		return false, err
	}
	return true, nil
}

//GraphQLDelete Eliminar
func GraphQLDelete(p graphql.ResolveParams, orm *xorm.Engine, data Table) (bool, error) {
	id, _ := p.Args["id"].(int)
	orm.Id(id).Get(data)
	reflections.SetField(data, "State", -1)
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		return false, err
	}
	_, err = orm.Id(id).Update(data)
	if err != nil {
		return false, err
	}
	return true, nil
}
