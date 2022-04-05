package model

type Hotel struct {
	HotelID      int    `json:"hotel_id,omitempty"`
	HotelName    string `json:"hotel_name,omitempty"`
	HotelCity    string `json:"hotel_city,omitempty"`
	HotelAddress string `json:"hotel_Addres,omitempty"`
	HotelPhone   string `json:"hotel_phone,omitempty"`
	Rooms        []Room `json:"rooms,omitempty"`
}
type HotelResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Hotel  `json:"data"`
}

type HotelsResponses struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Hotel `json:"data"`
}
