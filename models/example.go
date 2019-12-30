package models

import "github.com/jinzhu/gorm"

type Example struct {
	gorm.Model
	Data string `json:"data"`
}
