package data

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UserID      int
	Name        string
	Price       int
	Description string
	Location    string
	UserName    string
	Pictures    []Pictures
}

type Pictures struct {
	gorm.Model
	URL    string
	RoomID uint
}