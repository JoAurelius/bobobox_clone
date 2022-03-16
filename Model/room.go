package model

type Room struct {
	RoomID     int      `json:"room_id"`
	RoomType   RoomType `json:"room_type_id"`
	RoomNumber string   `json:"room_number"`
	RoomStatus string   `json:"room_status"`
}
