package data

import (
	"StayApp-API/features/rooms"
	"errors"

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

// SelectAllRoomUser implements rooms.RoomData
func (rm *roomQuery) SelectAllRoomUser(userid int, limit int, offset int) ([]rooms.Core, error) {
	// nameSearch := "%" + name + "%"
	var roomsModel []Room
	tx := rm.db.Limit(limit).Offset(offset).Where("user_id = ?", userid).Find(&roomsModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	roomsCoreAll := ListModelToCore(roomsModel)
	return roomsCoreAll, nil
}

// Delete implements rooms.RoomData
func (rm *roomQuery) Delete(userid int, id int) error {
	tx := rm.db.Where("user_id = ?", userid).Delete(&Room{}, id)
	if tx.RowsAffected < 1 {
		return errors.New("no data deleted")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements rooms.RoomData
func (rm *roomQuery) Update(userId uint, input rooms.Core, id uint) error {
	data := CoreToRoom(input)
	tx := rm.db.Model(&Room{}).Where("id = ? AND user_id = ?", id, userId).Updates(&data)
	if tx.RowsAffected < 1 {
		return errors.New("room no updated")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectOne implements rooms.RoomData
func (rm *roomQuery) SelectOne(id uint) (rooms.Core, error) {
	tmp := Room{}
	tx := rm.db.Where("id = ?", id).First(&tmp)
	if tx.RowsAffected < 1 {
		return rooms.Core{}, errors.New("room not found")
	}
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	rm.db.Raw("SELECT AVG(rating) FROM feedbacks Where room_id = ?", id).Scan(&tmp.RatingRoom)
	return RoomToCore(tmp), nil
}

// SelectAll implements rooms.RoomData
func (rm *roomQuery) SelectAll(limit int, offset int, name string) ([]rooms.Core, error) {
	nameSearch := "%" + name + "%"
	var roomsModel []Room
	tx := rm.db.Limit(limit).Offset(offset).Where("rooms.name LIKE ?", nameSearch).Select("rooms.id, rooms.name, rooms.price, rooms.pictures, rooms.availability, rooms.description, rooms.location, users.name AS user_name, AVG(feedbacks.rating) AS rating_room").Joins("JOIN users ON rooms.user_id = users.id").Joins("LEFT JOIN feedbacks ON rooms.id = feedbacks.room_id").Group("rooms.id").Find(&roomsModel)
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
