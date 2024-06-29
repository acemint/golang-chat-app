package main

import (
	"flag"

	"chat-app/internal/configuration"
	controller "chat-app/internal/controller/gin"
	"chat-app/internal/logging"
	repository "chat-app/internal/repository/postgres"
	"chat-app/internal/service"
)

func main() {
	flagProfile := flag.String("p", "", "Determines the configuration file to be used")
	flag.Parse()

	if err := configuration.Initialize(*flagProfile); err != nil {
		logging.Log.Error(err)
		panic(err)
	}

	if err := repository.InitializeConnection(); err != nil {
		logging.Log.Error(err)
		panic(err)
	}
	repository.InitializeRepositories(repository.DB)

	service.InitializeService(repository.DB, repository.MemberRepository)

	controller.InitializeGinServer()
	controller.InitializeRoutes(controller.Server, service.MemberService)
	if err := controller.StartGinServer(); err != nil {
		logging.Log.Error(err)
		panic(err)
	}

}
