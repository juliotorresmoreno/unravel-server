package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/unravel-server/graphql/users"
)

var schema graphql.Schema

//ExecuteQuery Ejecuta las consultas
func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func concat(store map[string]*graphql.Field, append map[string]*graphql.Field) map[string]*graphql.Field {
	for i, v := range append {
		store[i] = v
	}
	return store
}

func init() {
	var query = make(map[string]*graphql.Field, 0)
	var mutation = make(map[string]*graphql.Field, 0)
	query = concat(query, users.GetData)
	mutation = concat(mutation, users.SetData)

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: query,
	})

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutation,
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
