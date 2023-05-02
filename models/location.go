package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	Id        uint32  `json:"id" gorm:"primary_key"`
	CityId    uint32  `fake:"{number:1,10}"`
	Latitude  float32 `gorm:"type:numeric(10,2)" fake:"{latitude}"`
	Longitude float32 `gorm:"type:numeric(10,2)" fake:"{longitude}"`
	Address   string  `gorm:"type:varchar(255)" fake:"{streetname}"`
	City      City
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `sql:"index"`
}
