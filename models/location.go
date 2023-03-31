package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	CityID    uint    `fake:"{number:1,10}"`
	Latitude  float32 `gorm:"type:numeric(10,2)" fake:"{latitude}"`
	Longitude float32 `gorm:"type:numeric(10,2)" fake:"{longitude}"`
	Address   string  `gorm:"type:varchar(255)" fake:"{streetname}"`
	City      City
}
