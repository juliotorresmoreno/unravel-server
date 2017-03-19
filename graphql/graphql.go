package graphql

import (
	"log"
	"net/http"

	"../models"
	"../ws"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

//GetHandler api graphql
func GetHandler() func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "kjgh guhg ftf", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	return func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		h.ServeHTTP(w, r)
	}
}
