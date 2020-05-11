package services

import (
	"github.com/globalsign/mgo/bson"
	"my-todo-api/repository"
	"my-todo-api/schema"
)

type TaskService interface {
	ListTasks(taskData schema.TaskListFilter) ([]schema.Task, error)
	GetTaskById(id int) (schema.Task, error)
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

func (m taskService) ListTasks(taskData schema.TaskListFilter) ([]schema.Task, error) {
	var filter schema.FilterParam
	filter.Condition = bson.M{}
	if taskData.Type > 0 {
		filter.Condition["type"] = taskData.Type
	}

	return m.taskRepo.List(filter)
}

func (m taskService) GetTaskById(id int) (schema.Task, error) {
	var result schema.Task
	var filter = schema.FilterParam{
		Condition: bson.M{"id": id},
	}

	tasks, err := m.taskRepo.List(filter)
	if err != nil {
		return result, err
	}

	result = tasks[0]

	return result, nil
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