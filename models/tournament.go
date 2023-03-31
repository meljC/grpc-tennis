package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255)" fake:"{beername}"`
	DateStart time.Time `gorm:"type:DATE" fake:"{date}"`
	DateEnd   time.Time `gorm:"type:DATE" fake:"{date}"`
	// city
}
