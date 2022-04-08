package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Host   string
	Port   string
	Shell  string
	config *confStruct
)

type confStruct struct {
	Host  string `json : "Host"`
	Port  string `json : "Port"`
	Shell string `json : "Shell"`
}

func readConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config")
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Host = config.Host
	Port = config.Port
	Shell = config.Shell
	return nil
}
