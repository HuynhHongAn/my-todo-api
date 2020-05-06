package services

import (
	"github.com/globalsign/mgo/bson"
	"my-todo-api/repository"
	"my-todo-api/schema"
)

type TaskService interface {
	List() ([]schema.Task, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService() TaskService{
	return &taskService{
		taskRepo: repository.NewTaskRepository(),
	}
}

func (m taskService) List() ([]schema.Task, error) {
	return m.taskRepo.List(bson.M{})
}