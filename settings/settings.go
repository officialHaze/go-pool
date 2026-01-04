package settings

import (
	"fmt"
)

func Generate() error {
	conf, err := readSettingsConf()
	if err != nil {
		return fmt.Errorf("error generating settings: %v", err)
	}

	MySettings = &Settings{
		ENV_FILE_NAME: conf.Env_File_Name,
		SERVER_PORT:   conf.Server_Port,
	}

	return nil
}

type Settings struct {
	ENV_FILE_NAME string
	SERVER_PORT   int
}
