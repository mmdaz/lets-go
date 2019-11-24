package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	//Profile  Profile `gorm:"foreignkey:UserID"`
}
