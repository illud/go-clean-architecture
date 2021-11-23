package tasks

import (
	"fmt"

	"github.com/saturnavt/clean-architecture/db"
	prisma "github.com/saturnavt/clean-architecture/infraestructure/databases/db"
)

func Create(title string, desc string) string {
	// create a Taks
	_, err := prisma.Client.Tasks.CreateOne(
		db.Tasks.Title.Set(title),
		db.Tasks.Desc.Set(desc),
	).Exec(prisma.Context)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}

	return "Saved"
}

func FindAll() interface{} {
	// find a single tasks
	tasksResult, err := prisma.Client.Tasks.FindMany().Exec(prisma.Context)
	if err != nil {
		fmt.Println(err)
	}
	return tasksResult
}
