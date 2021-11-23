package tasks

import (
	tasksRepository "github.com/saturnavt/clean-architecture/infraestructure/repository/tasks"
	errors "github.com/saturnavt/clean-architecture/infraestructure/utils/errors"
)

func CreateTasks(title string, desc string) string {
	if title == "" || desc == "" {
		return errors.InvalidInfo
	}
	result := tasksRepository.Create(title, desc)
	if result == "Error" {
		return "Error"
	}
	return result
}

func GetTasks() interface{} {
	return tasksRepository.FindAll()
}
