package repository

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"my-todo-api/database"
	"my-todo-api/schema"
)

type TaskRepository interface{
	List(condition bson.M) ([]schema.Task, error)
}

type taskRepository struct {
	mgo database.MongoLayer
}

func NewTaskRepository () TaskRepository {
	return &taskRepository{
		mgo: database.NewMongoLayer("local", "tasks"),
	}
}

func (m taskRepository) List(condition bson.M) ([]schema.Task, error) {
	var result []schema.Task
	data, err := m.mgo.Find(condition)
	if err != nil {
		return nil, err
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
