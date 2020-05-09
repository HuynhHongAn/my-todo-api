package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-todo-api/routes"
)

func main() {
	fmt.Println("Hello world!")
	r := gin.Default()
	userGroup := r.Group("/")
	routes.RouteTask(userGroup)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}