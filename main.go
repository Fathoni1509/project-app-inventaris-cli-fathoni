package main

import (
	"context"
	"log"
	"project-app-inventaris-cli-fathoni/cmd"
	"project-app-inventaris-cli-fathoni/database"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/repository"
	"project-app-inventaris-cli-fathoni/service"
)

func main() {
	// init
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// init object
	repoInventaris := repository.NewRepoInventaris(db)
	serviceInventaris := service.NewServiceInventaris(&repoInventaris)
	handlerInventaris := handler.NewHandlerInventaris(&serviceInventaris)

	cmd.HomePage(handlerInventaris)
}