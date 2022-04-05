package model

type Promo struct {
	PromoCode       string  `json:"promo_code,omitempty"`
	PromoTitle      string  `json:"promo_title,omitempty"`
	PromoDesc       string  `json:"promo_desc,omitempty"`
	PromoPercentage float32 `json:"promo_percentage,omitempty"`
	PromoMax        string  `json:"promo_max,omitempty"`
	PromoCreated    string  `json:"promo_created,omitempty"`
	PromoEndDate    string  `json:"promo_end_date,omitempty"`
}

type PromoResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Promo  `json:"data"`
}

type PromosResponses struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Promo `json:"data"`
}
