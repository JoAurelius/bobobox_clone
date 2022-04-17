package model

type Hotel struct {
	HotelID      int    `json:"hotel_id,omitempty" gorm:"primaryKey"`
	HotelName    string `json:"hotel_name,omitempty"`
	HotelCity    string `json:"hotel_city,omitempty"`
	HotelAddress string `json:"hotel_address,omitempty"`
	HotelPhone   string `json:"hotel_phone,omitempty"`
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
