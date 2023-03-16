package handler

type RoomRequest struct {
	Name              string `json:"name" form:"name"`
	Price             int    `json:"price" form:"price"`
	Description       string `json:"description" form:"description"`
	Location          string `json:"location" form:"location"`
	MaxVisitors       int    `json:"max_visitors" form:"max_visitors"`
	SpecialAccess     string `json:"special_access" form:"special_access"`
	Bedroom           int    `json:"bedroom" form:"bedroom"`
	RoomTotal         int    `json:"room_total" form:"room_total"`
	Kitchen           string `json:"kitchen" form:"kitchen"`
	Wifi              string `json:"wifi" form:"wifi"`
	Garage            string `json:"garage" form:"garage"`
	ExcellentFeatures string `json:"excellent_features" form:"excellent_features"`
	CleaningFee       int    `json:"cleaning_fee" form:"cleaning_fee"`
	ServiceFee        int    `json:"service_fee" form:"service_fee"`
	BringAnimal       string `json:"bring_animal" form:"bring_animal"`
	CheckIn           string `json:"check_in" form:"check_in"`
	CheckOut          string `json:"check_out" form:"check_out"`
	TakePhoto         string `json:"take_photo" form:"take_photo"`
	Other             string `json:"other" form:"other"`
	UserID            uint   `json:"user_id" form:"user_id"`
	UserName          string `json:"user_name" form:"user_name"`
}
