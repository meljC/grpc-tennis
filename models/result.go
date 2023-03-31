package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	PlayerID uint `fake:"{number:1,10}"`
	MatchID  uint `fake:"{number:1,10}"`
	UserID   uint `fake:"{number:1,10}"`
	Points   int  `fake:"{number:1,200}"`
	Player   Player
	Match    Match
	User     User
}
