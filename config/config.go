package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var cnfg *Config

func ConfigIniti(pathConfig string) (*Config, error) {

	openFile, err := os.Open(pathConfig)

	if err != nil {
		log.Println("configuration error", err.Error())
		return nil, err
	}
	log.Println("Configuration successfully")

	defer openFile.Close()
	Val, errf := io.ReadAll(openFile)
	if errf != nil {
		log.Println("configuration error Reading File", errf.Error())
	}

	var output Config

	if err := json.Unmarshal(Val, &output); err != nil {
		log.Println("Unmarshal Error", err.Error())

	}
	cnfg = &output

	return cnfg, nil
}
func GetConf() *Config {
	return cnfg
}
