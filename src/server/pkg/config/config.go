package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Agents []string `json:"Agents"`
	Port string `json:"Port"`
  User string `json:"user"`
  Password string `json:"password"`
  DbHost string `json:"dbhost"`
}

func NewConfig(cfg_file string) (*Config, error) {
	var config *Config
	file, err := os.ReadFile(cfg_file)
	if err != nil {
		log.Println("Error reading config -> ", err)
		return nil,err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println("Error unmarshalling config -> ", err.Error())
		return nil,err
	}
	return config,err
}
