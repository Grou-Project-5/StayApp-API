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

// SelectAll implements rooms.RoomData
func (rm *roomQuery) SelectAll(limit int, offset int, name string) ([]rooms.Core, error) {
	nameSearch := "%" + name + "%"
	var roomsModel []Room
	tx := rm.db.Limit(limit).Offset(offset).Where("rooms.name LIKE ?", nameSearch).Select("rooms.id, rooms.name, rooms.price, rooms.description, rooms.location, users.name").Joins("JOIN users ON rooms.user_id = users.id").Find(&roomsModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	roomsCoreAll := ListModelToCore(roomsModel)
	return roomsCoreAll, nil
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
