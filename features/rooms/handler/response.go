package handler

import "StayApp-API/features/rooms"

type RoomResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
	UserID      uint    `json:"user_id" form:"user_id"`
	UserName    string `json:"user_name" form:"user_name"`
}

func CoreToGetAllRoomResp(data []rooms.Core) []RoomResponse {
	res := []RoomResponse{}
	for _, val := range data {
		res = append(res, CoreToGetRoomResp(val))
	}
	return res
}

func CoreToGetRoomResp(data rooms.Core) RoomResponse {
	return RoomResponse{
		ID:          uint(data.ID),
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Location:    data.Location,
		UserID:      data.UserID,
		UserName:    data.UserName,
	}
}
