package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	TcpAddr string
	Shell  string
	Platform string
	config *confStruct
)

type confStruct struct {
	TcpAddr string `json:"TCPAddr"`
	Shell string `json:"Shell"`
	Plaftorm string `json:"Platform"`
}

func ReadConfig() error {
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
	TcpAddr = config.TcpAddr
	Shell = config.Shell
	Platform = config.Plaftorm
	return nil
}
