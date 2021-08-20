package db

import (
	"errors"
	"fmt"
)

type Database interface {
	GetTask(id string) (Task, error)
	GetTasks() ([]Task, error)
	PutTask(id string, task Task) (Task, error)
}

type LocalDB struct {
	data map[string]Task
}

var data = map[string]Task{
	"1": {
		ID: "1",
		Title: "This is title 1",
	},
	"2": {
		ID: "2",
		Title: "This is title 2",
	},
	"3": {
		ID: "3",
		Title: "This is title 3",
	},
}

func NewLocalDB() *LocalDB {
	return &LocalDB{
		data: data,
	}
}

func (db *LocalDB) GetTask(id string) (Task, error) {
	if v, ok := db.data[id]; ok {
		return v, nil
	}
	return Task{}, fmt.Errorf("could not find task: %s", id)
}

func (db *LocalDB) GetTasks() ([]Task, error) {
	if len(db.data) == 0 {
		return []Task{}, errors.New("no data")
	}

	var tasks []Task
	for _, v := range db.data {
		tasks = append(tasks, v)
	}

	return tasks,nil
}

func (db *LocalDB) PutTask(id string, task Task) (Task, error) {
	if _, ok := db.data[id]; ok {
		data[id] = task
		return task, nil
	}

	return Task{}, fmt.Errorf("could not find task: %s", id)
}

