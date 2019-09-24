package main

import (
	"lets_go/gorm_practice/cmd"
	"lets_go/gorm_practice/pkg/log"
)

func main() {
	cmd.Main()
	log.Logger.Info("main started...")
}
