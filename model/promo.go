package model

type Promo struct {
	PromoCode       string  `json:"promo_code"`
	PromoTitle      string  `json:"promo_title"`
	PromoDesc       string  `json:"promo_desc"`
	PromoPercentage float32 `json:"promo_percentage"`
	PromoMax        string  `json:"promo_max"`
	PromoCreated    string  `json:"promo_created"`
	PromoEndDate    string  `json:"promo_end_date"`
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
