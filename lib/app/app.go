package app

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"

	"github.com/cnnrrss/gql-go-server/lib/db"
	"github.com/cnnrrss/gql-go-server/lib/gql"
)

func Run() {
	schemaData, err := gql.GetSchemaData()
	if err != nil {
		panic(err)
	}

	resolver := gql.NewResolver(db.NewLocalDB())

	schema, err := graphql.ParseSchema(string(schemaData), resolver)
	if err != nil {
		panic(err)
	}

	gqlHandler := gql.NewHandler(schema)

	http.HandleFunc("/gql", gqlHandler.ServerHTTP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

