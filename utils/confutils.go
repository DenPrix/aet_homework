package utils

import (
	"encoding/json"
	"log"
	"os"
)

var Config *Configuration

type Configuration struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Endpoint string
}

func InitConfig() {
	//Able to set config path in environment variable named ConfigPath
	configPath := "config.json"
	cp := os.Getenv("ConfigPath")
	if cp != "" {
		configPath = cp
	}
	file, err := os.Open(configPath)
	if err != nil {
		log.Println("ConfigPath is wrong... using default config file")
		file, err = os.Open("config.json")
		if err != nil {
			log.Println(err)
		}
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println(err)
	}
	Config = &configuration
}
