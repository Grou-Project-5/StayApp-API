package service

import (
	"StayApp-API/features/users"
	"StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"errors"
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

// Login implements users.UserService
func (us *userService) Login(email string, password string) (string, users.Core, error) {
	tmp, err := us.data.Login(email)
	if err != nil {
		return "", users.Core{}, err
	}
	// Compare password
	if err := helper.PassCompare(tmp.Password, password); err != nil {
		return "", users.Core{}, err
	}
	// Generate token
	token, err := middlewares.CreateToken(tmp.Id)
	if err != nil {
		return "", users.Core{}, err
	}
	return token, tmp, nil
}

// Profile implements users.UserService
func (us *userService) MyProfile(userID int) (users.Core, error) {
	tmp, err := us.data.MyProfile(userID)
	if err != nil {
		return users.Core{}, err
	}
	return tmp, nil
}

// Update implements users.UserService
func (us *userService) Update(userID int, updateUser users.Core, fileHeader *multipart.FileHeader) error {
	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/users")
		if err != nil {
			return err
		}
		updateUser.Pictures = uploadURL[0]
	}
	err := us.data.Update(userID, updateUser)
	if err != nil {
		return err
	}
	return nil
}

// ChangePassword implements users.UserService
func (us *userService) ChangePassword(userID int, oldPass string, newPass users.Core) error {
	if len(newPass.Password) < 3 {
		return errors.New("your password must be at least 3 characters")
	}
	// Get old password from database
	res, err := us.data.CheckPassword(userID)
	if err != nil {
		return err
	}
	// Compare old password and new password
	errCompare := helper.PassCompare(res.Password, oldPass)
	if errCompare != nil {
		return errCompare
	}
	// Bcrypt new password
	bcrypt, errBcrypt := helper.PassBcrypt(newPass.Password)
	if errBcrypt != nil {
		return errBcrypt
	}
	newPass.Password = bcrypt
	// Update password in database
	err = us.data.ChangePassword(userID, newPass)
	if err != nil {
		return err
	}
	return nil
}

// UserByID implements users.UserService
func (us *userService) UserByID(userID int) (users.Core, error) {
	tmp, err := us.data.UserByID(userID)
	if err != nil {
		return users.Core{}, err
	}
	return tmp, nil
}

// Delete implements users.UserService
func (us *userService) Delete(userID int) error {
	err := us.data.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}