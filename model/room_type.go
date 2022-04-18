package model

type RoomType struct {
	RoomTypeID      int    `json:"room_type_id,omitempty" gorm:"primaryKey"`
	RoomType        string `json:"room_type,omitempty"`
	RoomDescription string `json:"room_description,omitempty"`
	RoomPrice       int    `json:"room_price,omitempty"`
}

type RoomTypeResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    RoomType `json:"data"`
}

type RoomTypesResponses struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []RoomType `json:"data"`
}
