package service

import (
	"StayApp-API/features/users"
	"StayApp-API/middlewares"
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
func (*userService) ChangePassword(id int, oldPass string, newPass users.Core) error {
	panic("unimplemented")
}

// Delete implements users.UserService
func (*userService) Delete(id int) error {
	panic("unimplemented")
}

// UserByID implements users.UserService
func (*userService) UserByID(id int) (users.Core, error) {
	panic("unimplemented")
}
