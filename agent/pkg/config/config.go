package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TcpAddr  string `json:"TCPAddr"`
	Shell    string `json:"Shell"`
	Plaftorm string `json:"Platform"`
}

func NewConfig() (*Config, error) {
	var config *Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config", err)
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error unmarshalling config", err)
		return nil, err
	}
	return config, err
}
