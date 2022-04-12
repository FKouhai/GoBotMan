package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	c "agent/pkg/config"
)

func main() {
	config, err := c.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	listener(config)
}

func listener(config *c.Config) {
	listen, err := net.Listen("tcp", config.TcpAddr)
	if err != nil {
		log.Fatalln("Unable to start listener, port is in use")
	}
	defer listen.Close()
	fmt.Println("Listening on " + config.TcpAddr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection", err.Error())
		}
		go handleRequest(conn, config)

	}
}
func handleRequest(conn net.Conn, config *c.Config) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if ( config.Plaftorm == "Windows"){
			out, err := exec.Command(config.Shell, "/C", strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Println("unable to run command,", conn, "%s", err)
		}
		fmt.Fprintf(conn, "%s\n", out)
		}else {
			out, err := exec.Command(config.Shell, "-c", strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Println("unable to run command,", conn, "%s", err)
		}
		fmt.Fprintf(conn, "%s\n", out)
		}

	}
}

// TODO -> Include flags library to provide support to only allow connections from the central node to said port
