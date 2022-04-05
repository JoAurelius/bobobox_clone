package model

type Room struct {
	RoomID     int      `json:"room_id,omitempty"`
	RoomType   RoomType `json:"room_type_id,omitempty"`
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
