package models

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	// Drop table if exists
	db.Migrator().DropTable(&City{}, &Role{}, &User{}, &Player{}, &Location{}, &Match{}, &Result{}, &Tournament{})

	// Migrate the schema
	db.AutoMigrate(&City{}, &Role{}, &User{}, &Player{}, &Location{}, &Match{}, &Result{}, &Tournament{})

}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetLocalTime() string {
	return u.CreatedAt.Format("02.01.2006. 15:04:05")
}
