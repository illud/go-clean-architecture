package main

import (
	"github.com/gin-gonic/gin"
	tasksController "github.com/saturnavt/clean-architecture/controller/tasks"
)

func main() {
	router := gin.Default()
	//tasks
	router.GET("/tasks", tasksController.GetTasks)
	router.POST("/tasks", tasksController.CreateTasks)
	router.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
