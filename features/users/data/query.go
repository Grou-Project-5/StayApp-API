package data

import (
	"StayApp-API/features/users"
	"errors"

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

// Login implements users.UserData
func (uq *userQuery) Login(email string) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("email = ?", email).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("email not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

// Profile implements users.UserData
func (uq *userQuery) MyProfile(userID int) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("id = ?", userID).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("user not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

// Update implements users.UserData
func (uq *userQuery) Update(userID int, updateUser users.Core) error {
	data := CoreToUser(updateUser)
	tx := uq.db.Model(&User{}).Where("id = ?", userID).Updates(&data)
	if tx.RowsAffected < 1 {
		return errors.New("profile no updated")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// CheckPassword implements users.UserData
func (uq *userQuery) CheckPassword(userID int) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("id = ?", userID).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("user not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

// ChangePassword implements users.UserData
func (uq *userQuery) ChangePassword(userID int, newPass users.Core) error {
	tmp := CoreToUser(newPass)
	tx := uq.db.Model(&User{}).Where("id = ?", userID).Updates(&tmp)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UserByID implements users.UserData
func (uq *userQuery) UserByID(userID int) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("id = ?", userID).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("user not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

// Delete implements users.UserData
func (uq *userQuery) Delete(userID int) error {
	tx := uq.db.Delete(&User{}, userID)
	if tx.RowsAffected < 1 {
		return errors.New("no data deleted")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}