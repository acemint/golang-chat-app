package controller

import (
	"chat-app/internal/configuration"
	"chat-app/internal/logging"
	"errors"
	"strconv"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

const (
	DefaultPort uint16 = 8080
)

var Server *gin.Engine

func InitializeGinServer(profilerPath string) {
	Server = gin.Default()

	if profilerPath != "" {
		logging.Log.Debug("Added profiler to server in path " + profilerPath)
		pprof.Register(Server, profilerPath)
	}
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
