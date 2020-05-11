package routes

import (
	"github.com/gin-gonic/gin"
	"my-todo-api/handlers"
	"my-todo-api/middleware"
)

func RouteTask(r *gin.RouterGroup) {
	groupTask := r.Group("/tasks")
	groupTask.Use(middleware.SetupResponse())
	{
		taskHandler := handlers.TaskHandler{}
		groupTask.GET("/", taskHandler.ListTasks)
		groupTask.GET("/:id", taskHandler.GetTaskById)
		groupTask.POST("/add", taskHandler.AddTask)
		groupTask.POST("/edit/:id", taskHandler.EditTask)
		groupTask.POST("/remove/:id", taskHandler.RemoveTask)
	}
}
