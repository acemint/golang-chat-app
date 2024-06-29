package controller

import (
	"chat-app/internal/configuration"
	"chat-app/internal/logging"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPort uint16 = 8080
)

var Server *gin.Engine

func InitializeGinServer() {
	Server = gin.Default()
}

func StartGinServer() error {
	if Server == nil {
		err := errors.New("need to initialize the server before starting the server")
		logging.Log.Error(err)
		return err
	}
	if err := Server.Run(":" + getPort()); err != nil {
		logging.Log.Error(err)
		return err
	}
	return nil
}

func getPort() string {
	var port uint16
	if configuration.AppConfiguration.Server.Port == 0 {
		logging.Log.Infof("Defaulting to port %d as properties not set", DefaultPort)
		port = DefaultPort
	} else {
		logging.Log.Infof("Obtained port %d from properties", port)
		port = configuration.AppConfiguration.Server.Port
	}
	return strconv.Itoa(int(port))
}
