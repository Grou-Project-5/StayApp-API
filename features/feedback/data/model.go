package data

import (
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	RoomID int
	UserID int
	Rating int
	Comment int
}
