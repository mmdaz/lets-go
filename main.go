package main

import (
	"lets_go/config"
	"lets_go/log"
)

func main() {
	config.Initialize("config/test.yaml")
	log.Initialize(config.Conf.Log.Level)
	log.Logger.Info("main started...")

}
