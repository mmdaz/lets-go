package model

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
}
