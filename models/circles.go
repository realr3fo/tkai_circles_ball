package models

import "github.com/jinzhu/gorm"

type Logger struct {
	gorm.Model
	AccessedTime string `json:"accessedTime"`
	Username string `json:"username"`
	AccessedFunction string `json:"accessedFunction"`
}
