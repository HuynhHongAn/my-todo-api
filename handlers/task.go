package handlers

import (
	"github.com/gin-gonic/gin"
	"my-todo-api/common"
	"my-todo-api/services"
)

type TaskHandler struct {

}

func (m TaskHandler) ListTasks (c *gin.Context) {
	taskService := services.NewTaskService()
	data, err := taskService.List()
	if err != nil {
		c.JSON(200, common.Format(0,err, nil))
		return
	}
	c.JSON(200, common.Format(1,nil, data))
}