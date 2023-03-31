package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Player1ID  uint      `fake:"{number:1,10}"`
	Player2ID  uint      `fake:"{number:1,10}"`
	StartTime  time.Time `fake:"{date}"`
	LocationID uint      `fake:"{number:1,10}"`
	Location   Location
}
