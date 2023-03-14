package handler

type MiniResponse struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Pictures string `json:"pictures" form:"pictures"`
}

type UserResponse struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
	Pictures string `json:"pictures" form:"pictures"`
}