package gql

import "github.com/graph-gophers/graphql-go"

type TaskArgs struct {
	ID graphql.ID
}

type UpdateTaskArgs struct {
	ID graphql.ID
	Task struct {
		Title graphql.NullString
		Description graphql.NullString
	}
}

