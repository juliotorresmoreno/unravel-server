package users

import (
	"github.com/graphql-go/graphql"
	"github.com/juliotorresmoreno/unravel-server/crud"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/models"
)

var tipos = map[string]graphql.Type{
	"id":        graphql.Int,
	"nombres":   graphql.String,
	"apellidos": graphql.String,
	"fullName":  graphql.String,
	"email":     graphql.String,
	"usuario":   graphql.String,
	"passwd":    graphql.String,
	"recovery":  graphql.String,
	"tipo":      graphql.String,
	"code":      graphql.String,
	"create_at": graphql.String,
	"update_at": graphql.String,
}

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: tipos["id"],
		},
		"nombres": &graphql.Field{
			Type: tipos["nombres"],
		},
		"apellidos": &graphql.Field{
			Type: tipos["apellidos"],
		},
		"fullName": &graphql.Field{
			Type: tipos["fullName"],
		},
		"email": &graphql.Field{
			Type: tipos["email"],
		},
		"usuario": &graphql.Field{
			Type: tipos["usuario"],
		},
		"passwd": &graphql.Field{
			Type: tipos["passwd"],
		},
		"recovery": &graphql.Field{
			Type: tipos["recovery"],
		},
		"tipo": &graphql.Field{
			Type: tipos["tipo"],
		},
		"code": &graphql.Field{
			Type: tipos["code"],
		},
		"created_at": &graphql.Field{
			Type: tipos["created_at"],
		},
		"updated_at": &graphql.Field{
			Type: tipos["updated_at"],
		},
	},
})

//GetData Obtiene los datos
var GetData = graphql.Fields{
	"categoryList": &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "List of category",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var orm = db.GetXORM()
			defer orm.Close()
			data := make([]models.Category, 0)
			err := crud.GraphQLGet(params, orm, &data)
			return data, err
		},
	},
}

//SetData Establece los datos
var SetData = graphql.Fields{
	"createCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Create new category",
		Args: graphql.FieldConfigArgument{
			"nombres": &graphql.ArgumentConfig{
				Type: tipos["nombres"],
			},
			"apellidos": &graphql.ArgumentConfig{
				Type: tipos["apellidos"],
			},
			"fullName": &graphql.ArgumentConfig{
				Type: tipos["fullName"],
			},
			"email": &graphql.ArgumentConfig{
				Type: tipos["email"],
			},
			"usuario": &graphql.ArgumentConfig{
				Type: tipos["usuario"],
			},
			"passwd": &graphql.ArgumentConfig{
				Type: tipos["passwd"],
			},
			"recovery": &graphql.ArgumentConfig{
				Type: tipos["recovery"],
			},
			"tipo": &graphql.ArgumentConfig{
				Type: tipos["tipo"],
			},
			"code": &graphql.ArgumentConfig{
				Type: tipos["code"],
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var orm = db.GetXORM()
			defer orm.Close()
			data := models.User{}
			_, err := crud.GraphQLPut(params, orm, &data)
			return data, err
		},
	},
	"updateCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Update existing category, mark it done or not done",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["id"]),
			},
			"nombres": &graphql.ArgumentConfig{
				Type: tipos["nombres"],
			},
			"apellidos": &graphql.ArgumentConfig{
				Type: tipos["apellidos"],
			},
			"fullName": &graphql.ArgumentConfig{
				Type: tipos["fullName"],
			},
			"email": &graphql.ArgumentConfig{
				Type: tipos["email"],
			},
			"usuario": &graphql.ArgumentConfig{
				Type: tipos["usuario"],
			},
			"passwd": &graphql.ArgumentConfig{
				Type: tipos["passwd"],
			},
			"recovery": &graphql.ArgumentConfig{
				Type: tipos["recovery"],
			},
			"tipo": &graphql.ArgumentConfig{
				Type: tipos["tipo"],
			},
			"code": &graphql.ArgumentConfig{
				Type: tipos["code"],
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var orm = db.GetXORM()
			defer orm.Close()
			data := models.User{}
			crud.GraphQLPost(params, orm, &data)
			return data, nil
		},
	},
	"deleteCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Delete existing category",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["id"]),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var orm = db.GetXORM()
			defer orm.Close()
			data := models.User{}
			crud.GraphQLDelete(params, orm, &data)
			return data, nil
		},
	},
}
