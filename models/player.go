package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	FirstName string    `gorm:"type:varchar(255)" fake:"{noun}" json:"firstname"`
	LastName  string    `gorm:"type:varchar(255)" fake:"{lunch}" json:"lastname"`
	BirthDate time.Time `gorm:"type:DATE" fake:"{year}-{month}-{day}" format:"2006-01-02" json:"birthdate"`
	CityId    uint      `fake:"{number:1,10}" json:"cityid"`
	// city za turnir?
}
