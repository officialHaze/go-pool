package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func readSettingsConf() (*SettingsConf, error) {
	filepath := path.Join("settings", "settings.conf.jsonc")

	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %v", filepath, err)
	}

	dec := json.NewDecoder(f)

	conf := &SettingsConf{}
	if err := dec.Decode(conf); err != nil {
		return nil, fmt.Errorf("error while decoding settings configuration: %v", err)
	}

	return conf, nil
}

type SettingsConf struct {
	Env_File_Name string `json:"env_file_name"` // which env file to use
	Server_Port   int    `json:"server_port"`
}
