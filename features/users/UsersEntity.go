package users

import "mime/multipart"

type Core struct {
	ID       int
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3"`
	Address  string `validate:"required"`
	Phone    string `validate:"required"`
	Gender   string `validate:"required"`
	Pictures string
}

type UserService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
	MyProfile(id int) (Core, error)
	Update(id int, updateUser Core, file *multipart.FileHeader) error
	ChangePassword(id int, oldPass string, newPass Core) error
	UserByID(id int) (Core, error)
	Delete(id int) error
}

type UserData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
	MyProfile(id int) (Core, error)
	Update(id int, updateUser Core) error
	CheckPassword(id int) (Core, error)
	ChangePassword(id int, newPass Core) error
	UserByID(id int) (Core, error)
	Delete(id int) error
}
