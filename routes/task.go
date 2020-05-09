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
		groupTask.POST("/add", taskHandler.AddTask)
		groupTask.POST("/edit/:id", taskHandler.EditTask)
		groupTask.POST("/remove/:id", taskHandler.RemoveTask)
	}
}
