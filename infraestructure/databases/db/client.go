package client

import (
	"fmt"

	"github.com/saturnavt/clean-architecture/db"
	"golang.org/x/net/context"
)

var Client = DB()

func DB() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}

	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return client
}

var Context = ContextService()

func ContextService() context.Context {
	ctx := context.Background()
	return ctx
}
