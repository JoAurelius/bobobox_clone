package model

type Room struct {
	RoomID     int      `json:"room_id,omitempty"`
	HotelID    int      `json:"-"`
	Hotel      Hotel    `json:"hotel,omitempty" gorm:"foreignKey:HotelID;references:HotelID"`
	RoomTypeID int      `json:"-"`
	RoomType   RoomType `json:"room_type,omitempty" gorm:"foreignKey:RoomTypeID;references:RoomTypeID"`
	RoomNumber string   `json:"room_number,omitempty"`
	RoomStatus string   `json:"room_status,omitempty"`
}
type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}
