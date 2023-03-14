package data

import (
	"StayApp-API/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements users.UserData
func (uq *userQuery) Register(newUser users.Core) error {
	data := CoreToUser(newUser)
	tx := uq.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// ChangePassword implements users.UserData
func (*userQuery) ChangePassword(id int, newPass users.Core) error {
	panic("unimplemented")
}

// CheckPassword implements users.UserData
func (*userQuery) CheckPassword(id int) (users.Core, error) {
	panic("unimplemented")
}

// Delete implements users.UserData
func (*userQuery) Delete(id int) error {
	panic("unimplemented")
}

// Login implements users.UserData
func (*userQuery) Login(email string) (users.Core, error) {
	panic("unimplemented")
}

// MyProfile implements users.UserData
func (*userQuery) MyProfile(id int) (users.Core, error) {
	panic("unimplemented")
}

// Update implements users.UserData
func (*userQuery) Update(id int, updateUser users.Core) error {
	panic("unimplemented")
}

// UserByID implements users.UserData
func (*userQuery) UserByID(id int) (users.Core, error) {
	panic("unimplemented")
}
