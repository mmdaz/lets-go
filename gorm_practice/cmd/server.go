package cmd

import (
	"lets_go/gorm_practice/internal/postgres"
	"lets_go/gorm_practice/pkg/config"
	"lets_go/gorm_practice/pkg/log"
)

func Main (){
	config.Initialize("gorm_practice/pkg/config/test.yaml")
	log.Initialize(config.Conf.Log.Level)
	_ = postgres.GetPostgresDB()
}