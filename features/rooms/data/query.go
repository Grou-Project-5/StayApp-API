package data

import (
	"StayApp-API/features/rooms"

	"gorm.io/gorm"
)

type roomQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) rooms.RoomData {
	return &roomQuery{
		db: db,
	}
}

// Insert implements rooms.RoomData
func (rm *roomQuery) Insert(input rooms.Core) error {
	data := CoreToRoom(input)
	tx := rm.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
