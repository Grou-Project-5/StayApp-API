package handler

import (
	"StayApp-API/features/rooms"
)

type RoomRequest struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
	UserID      int    `json:"user_id" form:"user_id"`
	UserName    string `json:"user_name" form:"user_name"`
}

func requestToCore(data interface{}) *rooms.Core {
	res := rooms.Core{}

	switch data.(type) {
	case RoomRequest:
		cnv := data.(RoomRequest)
		res.Name = cnv.Name
		res.Price = cnv.Price
		res.Description = cnv.Description
		res.Location = cnv.Location
		res.UserID = cnv.UserID
		res.UserName = cnv.UserName
	default:
		return nil
	}
	return &res
}
