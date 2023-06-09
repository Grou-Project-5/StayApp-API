package data

import (
	modelFeedback "StayApp-API/features/feedback/data"
	modelReservation "StayApp-API/features/reservations/data"
	"StayApp-API/features/rooms"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UserID            uint
	Name              string
	Price             int
	Description       string
	Location          string
	Availability      string `gorm:"default:Available"`
	UserName          string
	MaxVisitors       int
	SpecialAccess     string
	Bedroom           int
	RoomTotal         int
	Kitchen           string
	Wifi              string
	Garage            string
	ExcellentFeatures string
	CleaningFee       int
	ServiceFee        int
	BringAnimal       string
	CheckIn           string
	CheckOut          string
	TakePhoto         string
	Other             string
	UserPhone         string
	Pictures          string
	RatingRoom        float64
	Feedback          []modelFeedback.Feedback
	Reservation       []modelReservation.Reservation
}

func CoreToRoom(data rooms.Core) Room {
	return Room{
		Model:             gorm.Model{ID: data.ID},
		UserID:            data.UserID,
		Name:              data.Name,
		Price:             data.Price,
		Description:       data.Description,
		Location:          data.Location,
		UserName:          data.UserName,
		MaxVisitors:       data.MaxVisitors,
		SpecialAccess:     data.SpecialAccess,
		Bedroom:           data.Bedroom,
		RoomTotal:         data.RoomTotal,
		Kitchen:           data.Kitchen,
		Wifi:              data.Wifi,
		Garage:            data.Garage,
		ExcellentFeatures: data.ExcellentFeatures,
		CleaningFee:       data.CleaningFee,
		ServiceFee:        data.ServiceFee,
		BringAnimal:       data.BringAnimal,
		CheckIn:           data.CheckIn,
		CheckOut:          data.CheckOut,
		TakePhoto:         data.TakePhoto,
		Other:             data.Other,
		UserPhone:         data.UserPhone,
		Pictures:          data.Pictures,
		RatingRoom:        data.RatingRoom,
		Availability:      data.Availability,
	}
}

func RoomToCore(data Room) rooms.Core {
	return rooms.Core{
		ID:                data.ID,
		Name:              data.Name,
		Price:             data.Price,
		Description:       data.Description,
		Location:          data.Location,
		UserID:            data.UserID,
		UserName:          data.UserName,
		MaxVisitors:       data.MaxVisitors,
		SpecialAccess:     data.SpecialAccess,
		Bedroom:           data.Bedroom,
		RoomTotal:         data.RoomTotal,
		Kitchen:           data.Kitchen,
		Wifi:              data.Wifi,
		Garage:            data.Garage,
		ExcellentFeatures: data.ExcellentFeatures,
		CleaningFee:       data.CleaningFee,
		ServiceFee:        data.ServiceFee,
		BringAnimal:       data.BringAnimal,
		CheckIn:           data.CheckIn,
		CheckOut:          data.CheckOut,
		TakePhoto:         data.TakePhoto,
		Other:             data.Other,
		Pictures:          data.Pictures,
		RatingRoom:        data.RatingRoom,
		Availability:      data.Availability,
	}
}

func ListModelToCore(dataModel []Room) []rooms.Core {
	var dataCore []rooms.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, RoomToCore(v))
	}
	return dataCore
}
