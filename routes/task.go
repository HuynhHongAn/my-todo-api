package routes

import (
	"github.com/gin-gonic/gin"
	"my-todo-api/handlers"
)

func RouteTask(r *gin.RouterGroup) {
	groupTask := r.Group("/tasks")
	{
		taskHandler := handlers.TaskHandler{}
		groupTask.GET("/", taskHandler.ListTasks)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
