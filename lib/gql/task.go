package gql

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/cnnrrss/gql-go-server/lib/db"
)


type TaskResolver struct {
	id graphql.ID
	task db.Task
}

func (t *TaskResolver) ID() graphql.ID {
	return t.id
}

func (t *TaskResolver) Title() string {
	return t.task.Title
}

func (t *TaskResolver) Description() *string {
	s := "desc"
	return &s
}
