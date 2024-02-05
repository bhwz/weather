package main

import (
	"github.com/bhwz/weather/weather-server/webapi"
	"os"
	"strconv"
)

func main() {
	var serverConfig webapi.ServerConfig
	var err error

	// Parse env.
	if serverConfig.Debug, err = strconv.ParseBool(os.Getenv("DEBUG")); err != nil {
		serverConfig.Debug = false
	}
	serverConfig.HttpPort = os.Getenv("HTTP_PORT")
	if serverConfig.HttpPort == "" {
		serverConfig.HttpPort = "8080"
	}

	// Start core web api service.
	webapi.Start(serverConfig)
}
