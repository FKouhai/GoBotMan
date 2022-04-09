package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Listening on " + comms_data)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command(config.Shell, "-c", strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Println("unable to run command,", conn, "%s\n", err)
		}
		fmt.Fprintf(conn, "%s\n", out)
	}
}

// TODO -> Include flags library to provide support to only allow connections from the central node to said port
