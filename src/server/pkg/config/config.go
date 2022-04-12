package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type confStruct struct {
	Network string `json:"Network"`
	Port string `json:"Port"`
}

func NewConfig() (*Config, error) {
	var config *Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config", err)
		return nil,err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config", err.Error())
		return nil,err
	}
	return config,err
}
