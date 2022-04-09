package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"agent/pkg/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	listener(config.TcpAddr)
}

func listener(comms_data string) {
	listen, err := net.Listen("tcp", comms_data)

	if err != nil {
		log.Fatalln("Unable to start listener, port is in use")
	}
	defer listen.Close()
	fmt.Println("Listening on " + comms_data)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection", err.Error())
		}
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if (config.Platform == "Windows"){
			out, err := exec.Command(config.Shell, "/C", strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Println("unable to run command,", conn, "%s", err)
		}
		fmt.Fprintln(conn, "%s", out)
		}else {
			out, err := exec.Command(config.Shell, "-c", strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Println("unable to run command,", conn, "%s", err)
		}
		fmt.Fprintln(conn, "%s", out)
		}

	}
}

// TODO -> Include flags library to provide support to only allow connections from the central node to said port
