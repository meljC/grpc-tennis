package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255)" fake:"{firstname}"`
	LastName  string `gorm:"type:varchar(255)" fake:"{lastname}"`
	Email     string `gorm:"unique;size:255" fake:"{email}"`
	Password  string `gorm:"type:varchar(255)" fake:"{password}"`
	RoleID    uint   `fake:"{number:1,2}"`
	Role      Role
}
