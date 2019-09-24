package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lets_go/gorm_practice/pkg/config"
	"lets_go/gorm_practice/pkg/log"
	"lets_go/gorm_practice/pkg/model"
	"os"
	"sync"
)

type Database struct {
	PostgresPool *gorm.DB
}

var (
	once             sync.Once
	postgresDatabase *Database
)

func GetPostgresDB() *Database {
	once.Do(func() {
		postgresDatabase = New()
	})
	return postgresDatabase
}

func New() *Database {
	pgPool := initialize()
	return &Database{
		PostgresPool: pgPool,
	}
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Profile{})
}

func initialize() *gorm.DB {
	var err error

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
		config.Conf.Postgres.Host, config.Conf.Postgres.Port, config.Conf.Postgres.User, config.Conf.Postgres.DB, config.Conf.Postgres.Pass))
	if err != nil {
		log.Logger.Error("Unable to create connection pool", "error", err)
		os.Exit(1)
	}

	db.LogMode(true)
	Migration(db)

	return db
}
