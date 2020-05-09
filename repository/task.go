package repository

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"my-todo-api/database"
	"my-todo-api/schema"
)

type TaskRepository interface{
	List(filterParam schema.FilterParam) ([]schema.Task, error)
	AddTask(task schema.Task) error
	EditTask(taskId int, task schema.Task) error
	RemoveTask(id int) error
}

type taskRepository struct {
	mgo database.MongoLayer
}

func NewTaskRepository () TaskRepository {
	return &taskRepository{
		mgo: database.NewMongoLayer("local", "tasks"),
	}
}

func (m taskRepository) List(filterParam schema.FilterParam) ([]schema.Task, error) {
	var result []schema.Task
	data, err := m.mgo.Find(filterParam)
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

func (m taskRepository) AddTask(task schema.Task) error {
	var lastId int = 1
	var filterParam = schema.FilterParam{
		Condition: nil,
		Skip:      0,
		Limit:     0,
		Sort:      "-id",
	}
	data, err := m.List(filterParam)
	if err != nil {
		return err
	}

	if len(data) > 0{
		lastId = data[0].Id + 1
	}
	task.Id = lastId

	return m.mgo.Insert(task)
}

func (m taskRepository) EditTask(taskId int, task schema.Task) error {
	condition := bson.M{
		"id": taskId,
	}
	var dataUpdate = bson.M{}
	if task.Title != "" {
		dataUpdate["title"] = task.Title
	}
	if task.Description != "" {
		dataUpdate["description"] = task.Description
	}

	return m.mgo.Update(condition, bson.M{"$set": dataUpdate})
}

func (m taskRepository) RemoveTask(id int) error {
	fmt.Println(id)
	condition := bson.M{
		"id": id,
	}
	return m.mgo.Remove(condition)
}