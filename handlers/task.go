package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-todo-api/common"
	"my-todo-api/schema"
	"my-todo-api/services"
	"strconv"
)

type TaskHandler struct {

}

func (m TaskHandler) ListTasks(c *gin.Context) {
	taskService := services.NewTaskService()
	var taskData schema.TaskListFilter
	var typeStr = c.Query("type")

	if taskType, err := strconv.Atoi(typeStr); err == nil {
		taskData.Type = taskType
	}

	data, err := taskService.ListTasks(taskData)
	if err != nil {
		c.JSON(200, common.Format(0,err, gin.H{}))
		return
	}
	c.JSON(200, common.Format(1,nil, data))
}

func (m TaskHandler) GetTaskById(c *gin.Context) {
	var taskId int
	var err error
	taskService := services.NewTaskService()
	idStr := c.Param("id")
	taskId, err = strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}

	task, err := taskService.GetTaskById(taskId)
	if err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}

	c.JSON(200, common.Format(1, nil, task))
	return


}

func (m TaskHandler) AddTask(c *gin.Context) {
	var task schema.Task
	taskService := services.NewTaskService()
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(200, common.Format(0,err, nil))
		return
	}

	if err := taskService.AddTask(task); err != nil {
		c.JSON(200, common.Format(0,err, nil))
		return
	}

	c.JSON(200, common.Format(1,nil, nil))
}

func (m TaskHandler) EditTask(c *gin.Context) {
	var task schema.Task
	var taskId int
	var err error
	taskService := services.NewTaskService()

	idStr := c.Param("id")
	taskId, err = strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}

	if err := taskService.EditTask(taskId, task); err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}
	c.JSON(200, common.Format(1, nil, nil))

}

func (m TaskHandler) RemoveTask(c *gin.Context) {
	var taskId int
	var err error
	taskService := services.NewTaskService()

	idStr := c.Param("id")
	taskId, err = strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}
	fmt.Println("Id", taskId)

	if err := taskService.RemoveTask(taskId); err != nil {
		c.JSON(200, common.Format(0, err, nil))
		return
	}

	c.JSON(200, common.Format(1, nil, taskId))

}