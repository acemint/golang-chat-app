package configuration

import (
	"encoding/json"
	"os"
	"path/filepath"

	"chat-app/internal/logging"
)

var AppConfiguration *Configuration
var FolderPath = "properties"
var FileName = "app-configuration"
var FileType = ".json"

func Initialize(profile string) error {
	if AppConfiguration != nil {
		return nil
	}

	var confFileName string
	if profile == "" {
		logging.Log.Infof("Initializing configuration without specific profile set")
		confFileName = FileName + FileType
	} else {
		logging.Log.Infof("Initializing configuration with -profile=%s", profile)
		confFileName = FileName + "-" + profile + FileType
	}

	basePath, err := os.Getwd()
	if err != nil {
		return err
	}

	configurationFile, err := os.ReadFile(filepath.Join(basePath, FolderPath, confFileName))
	if err != nil {
		return err
	}

	AppConfiguration = new(Configuration)
	err = json.Unmarshal(configurationFile, &AppConfiguration)
	if err != nil {
		return err
	}
	return nil
}
