package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/abgameur/jarvis/internal/config"
	"github.com/abgameur/jarvis/internal/protocol"
)

func main() {
	var configDir = flag.String("config-dir", "examples/config", "path to config directory")
	var socketPath = flag.String("socket", "/tmp/jarvis.sock", "path to socket")

	flag.Parse()

	os.Remove(*socketPath)

	listener, err := net.Listen("unix", *socketPath)
	if err != nil {
		fmt.Println("error while listening", err)
		return
	}

	fmt.Println("listening on", *socketPath)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error while accepting connection", err)
			return
		}

		handleConnection(conn, *configDir)
	}
}

func handleConnection(conn net.Conn, configDir string) {
	defer conn.Close()

	var request protocol.Request
	if err := json.NewDecoder(conn).Decode(&request); err != nil {
		fmt.Println("error while decoding request", err)
		return
	}

	log.Println(request.Op)

	response := protocol.Response{}
	switch request.Op {
	case "list_services":
		loadedConfig, err := config.ReadConfig(configDir)
		if err != nil {
			response.Error = fmt.Sprintf("read config: %v", err)
			break
		}

		response.Services = loadedConfig.Services
	default:
		response.Error = fmt.Sprintf("unsupported operation %q", request.Op)
	}

	if err := json.NewEncoder(conn).Encode(response); err != nil {
		fmt.Println("error while encoding response", err)
	}
}
