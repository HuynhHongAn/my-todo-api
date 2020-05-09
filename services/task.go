package services

import (
	"my-todo-api/repository"
	"my-todo-api/schema"
)

type TaskService interface {
	List() ([]schema.Task, error)
	AddTask(task schema.Task) error
	EditTask(id int, taskData schema.Task) error
	RemoveTask(id int) error
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
	return m.taskRepo.List(schema.FilterParam{})
}

func (m taskService) AddTask(task schema.Task) error {
	return m.taskRepo.AddTask(task)
}

func (m taskService) EditTask(id int, taskData schema.Task) error {
	return m.taskRepo.EditTask(id, taskData)
}

func (m taskService) RemoveTask(id int) error {
	return m.taskRepo.RemoveTask(id)
}