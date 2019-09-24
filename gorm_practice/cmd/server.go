package cmd

import (
	"lets_go/gorm_practice/internal/postgres"
	"lets_go/gorm_practice/internal/server"
	"lets_go/gorm_practice/pkg/config"
	"lets_go/gorm_practice/pkg/log"
)

func initialize() {
	config.Initialize("gorm_practice/pkg/config/test.yaml")
	log.NewLog(config.Conf.Log.Level)
	_ = postgres.GetPostgresDB()

}

func Main() {
	initialize()
	server.New().Run()
}
