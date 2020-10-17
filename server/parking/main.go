package main

import (
	"parking/config"
	"parking/internal/database"
	"parking/internal/predictor"
	serverPac "parking/internal/server"
)

func main() {
	config.LoadConfig("config")
	confServer := config.Config.Server
	confDB := config.Config.Database
	configPredict := config.Config.GrpcApi

	dbService := database.NewDatabase(confDB.Host, confDB.Port, confDB.Database, confDB.User, confDB.Password)
	predictorService := predictor.NewPredictor(configPredict.Host, configPredict.Port)
	server := serverPac.NewServer(confServer.Host, confServer.Port, dbService, predictorService)

	server.Run()
}