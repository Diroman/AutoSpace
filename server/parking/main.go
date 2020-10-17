package main

import (
	"parking/config"
	serverPac "parking/server"
)

func main() {
	config.LoadConfig("config")
	conf := config.Config

	server := serverPac.NewServer(conf.Server.Host, conf.Server.Port)

	server.Run()
}