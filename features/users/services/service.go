package service

import (
	"StayApp-API/features/users"
	"StayApp-API/utils/helper"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	data users.UserData
	vld  *validator.Validate
}

func New(repo users.UserData) users.UserService {
	return &userService{
		data: repo,
		vld:  validator.New(),
	}
}

// Register implements users.UserService
func (us *userService) Register(newUser users.Core) error {
	// Check input validation
	errVld := us.vld.Struct(newUser)
	if errVld != nil {
		return errVld
	}
	// Bcrypt password before insert into database
	passBcrypt, errBcrypt := helper.PassBcrypt(newUser.Password)
	if errBcrypt != nil {
		return errBcrypt
	}
	newUser.Password = passBcrypt
	err := us.data.Register(newUser)
	if err != nil {
		return err
	}
	return nil
}

// ChangePassword implements users.UserService
func (*userService) ChangePassword(id int, oldPass string, newPass users.Core) error {
	panic("unimplemented")
}

// Delete implements users.UserService
func (*userService) Delete(id int) error {
	panic("unimplemented")
}

// Login implements users.UserService
func (*userService) Login(email string, password string) (string, users.Core, error) {
	panic("unimplemented")
}

// MyProfile implements users.UserService
func (*userService) MyProfile(id int) (users.Core, error) {
	panic("unimplemented")
}

// Update implements users.UserService
func (*userService) Update(id int, updateUser users.Core, file *multipart.FileHeader) error {
	panic("unimplemented")
}

// UserByID implements users.UserService
func (*userService) UserByID(id int) (users.Core, error) {
	panic("unimplemented")
}
