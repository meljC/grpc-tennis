package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)" fake:"{city}"`
}
