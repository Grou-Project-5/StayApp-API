package handler

type MiniResponse struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Pictures string `json:"pictures" form:"pictures"`
}