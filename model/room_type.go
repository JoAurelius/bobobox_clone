package model

type RoomType struct {
	RoomTypeID      int    `json:"room_type_id"`
	RoomType        string `json:"room_type"`
	RoomDescription string `json:"room_description"`
	RoomPrice       int    `json:"room_price"`
}
