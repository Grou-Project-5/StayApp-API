package handler

import (
	"StayApp-API/features/rooms"
)

type RoomResponse struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name" form:"name"`
	Price             int     `json:"price" form:"price"`
	Description       string  `json:"description" form:"description"`
	Location          string  `json:"location" form:"location"`
	MaxVisitors       int     `json:"max_visitors" form:"max_visitors"`
	SpecialAccess     string  `json:"special_access" form:"special_access"`
	Bedroom           int     `json:"bedroom" form:"bedroom"`
	RoomTotal         int     `json:"room_total" form:"room_total"`
	Kitchen           string  `json:"kitchen" form:"kitchen"`
	Wifi              string  `json:"wifi" form:"wifi"`
	Garage            string  `json:"garage" form:"garage"`
	ExcellentFeatures string  `json:"excellent_features" form:"excellent_features"`
	CleaningFee       int     `json:"cleaning_fee" form:"cleaning_fee"`
	ServiceFee        int     `json:"service_fee" form:"service_fee"`
	BringAnimal       string  `json:"bring_animal" form:"bring_animal"`
	CheckIn           string  `json:"check_in" form:"check_in"`
	CheckOut          string  `json:"check_out" form:"check_out"`
	TakePhoto         string  `json:"take_photo" form:"take_photo"`
	Other             string  `json:"other" form:"other"`
	UserPhone         string  `json:"user_phone" form:"user_phone"`
	UserID            uint    `json:"user_id" form:"user_id"`
	UserName          string  `json:"user_name" form:"user_name"`
	Pictures          string  `json:"pictures"`
	RatingRoom        float64 `json:"rating_room"`
	Availability      string  `json:"availability" form:"availability"`
}

type AllRoomResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name" form:"name"`
	Price        int     `json:"price" form:"price"`
	Description  string  `json:"description" form:"description"`
	Location     string  `json:"location" form:"location"`
	UserName     string  `json:"user_name" form:"user_name"`
	Pictures     string  `json:"pictures"`
	RatingRoom   float64 `json:"rating_room"`
	Availability string  `json:"availability" form:"availability"`
}

func CoreToGetAllRoomResp(data []rooms.Core) []AllRoomResponse {
	res := []AllRoomResponse{}
	for _, val := range data {
		res = append(res, CoreToGetAllRoomRespB(val))
	}
	return res
}

func CoreToGetAllRoomRespB(data rooms.Core) AllRoomResponse {
	return AllRoomResponse{
		ID:           data.ID,
		Name:         data.Name,
		Price:        data.Price,
		Description:  data.Description,
		Location:     data.Location,
		UserName:     data.UserName,
		Pictures:     data.Pictures,
		RatingRoom:   data.RatingRoom,
		Availability: data.Availability,
	}
}

type GetOneRoomResponse struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name" form:"name"`
	Price             int     `json:"price" form:"price"`
	Description       string  `json:"description" form:"description"`
	Location          string  `json:"location" form:"location"`
	MaxVisitors       int     `json:"max_visitors" form:"max_visitors"`
	SpecialAccess     string  `json:"special_access" form:"special_access"`
	Bedroom           int     `json:"bedroom" form:"bedroom"`
	RoomTotal         int     `json:"room_total" form:"room_total"`
	Kitchen           string  `json:"kitchen" form:"kitchen"`
	Wifi              string  `json:"wifi" form:"wifi"`
	Garage            string  `json:"garage" form:"garage"`
	ExcellentFeatures string  `json:"excellent_features" form:"excellent_features"`
	CleaningFee       int     `json:"cleaning_fee" form:"cleaning_fee"`
	ServiceFee        int     `json:"service_fee" form:"service_fee"`
	BringAnimal       string  `json:"bring_animal" form:"bring_animal"`
	CheckIn           string  `json:"check_in" form:"check_in"`
	CheckOut          string  `json:"check_out" form:"check_out"`
	TakePhoto         string  `json:"take_photo" form:"take_photo"`
	Other             string  `json:"other" form:"other"`
	UserPhone         string  `json:"user_phone" form:"user_phone"`
	Pictures          string  `json:"pictures"`
	RatingRoom        float64 `json:"rating_room"`
	Availability      string  `json:"availability" form:"availability"`
}
