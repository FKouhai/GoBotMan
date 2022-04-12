package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TcpAddr string `json:"TCPAddr"`
	Shell string `json:"Shell"`
	Plaftorm string `json:"Platform"`
}

func ReadConfig() (*Config, error) {
	var config *Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config", err)
		return nil,err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return config, err
}
