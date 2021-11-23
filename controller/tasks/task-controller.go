package tasks

import (
	"github.com/gin-gonic/gin"
	tasksUseCase "github.com/saturnavt/clean-architecture/domain/useCase/tasks"
)

type Task struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func CreateTasks(c *gin.Context) {
	var requestBody Task
	c.ShouldBindJSON(&requestBody)

	c.JSON(200, gin.H{
		"message": tasksUseCase.CreateTasks(requestBody.Title, requestBody.Desc),
	})
}

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": tasksUseCase.GetTasks(),
	})
}
