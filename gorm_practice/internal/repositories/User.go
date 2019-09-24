package repositories

import (
	"lets_go/gorm_practice/internal/postgres"
	"lets_go/gorm_practice/pkg/log"
	"lets_go/gorm_practice/pkg/model"
)

type UserRepo struct {
	log      *log.Logger
	Database *postgres.Database
}

func (p UserRepo) Create(user *model.User) {
	// TODO I am so confusing :))))
}
