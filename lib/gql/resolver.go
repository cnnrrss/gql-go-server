package gql

import (
	"github.com/cnnrrss/gql-go-server/lib/db"
)

type Resolver struct {
	db.Database
}

func NewResolver(db db.Database) *Resolver {
	return &Resolver{db}
}

func (r *Resolver) Task(args TaskArgs) (*TaskResolver, error) {
	task, err := r.GetTask(string(args.ID))
	if err != nil {
		return nil, err
	}

	return &TaskResolver{
		id: args.ID,
		task: task,
	}, nil
}

func (r *Resolver) UpdateTask(args UpdateTaskArgs) (*TaskResolver, error) {
	var title, description string

	if args.Task.Title.Set {
		title = *args.Task.Title.Value
	}

	if args.Task.Description.Set {
		description = *args.Task.Description.Value
	}

	task, err := r.PutTask(string(args.ID), db.Task{
		ID: string(args.ID),
		Title: title,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	return &TaskResolver{
		id: args.ID,
		task: task,
	}, nil
}