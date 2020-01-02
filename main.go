package main

import (
	"lets_go/fast_http_example"
	"os"
	"os/signal"
	"syscall"

	//"lets_go/gorm_practice/cmd"
	"lets_go/gorm_practice/pkg/log"
)

func main() {
	//cmd.Main()
	fast_http_example.StartServer()
	log.Logger.Info("main started...")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
