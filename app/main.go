package main

import (
	"flag"

	"chat-app/internal/configuration"
	"chat-app/internal/logging"
	"chat-app/internal/repository"
)

func main() {
	flagProfile := flag.String("p", "", "Determines the configuration file to be used")
	flag.Parse()

	if err := configuration.Initialize(*flagProfile); err != nil {
		logging.Log.Error(err)
		panic(err)
	}

	if err := repository.InitializePostgreSqlConnection(); err != nil {
		logging.Log.Error(err)
		panic(err)
	}

	// TODO: Create connection using Gin. Some of the API will include: POST /member, GET /member, PUT /transaction, GET /transaction

}
