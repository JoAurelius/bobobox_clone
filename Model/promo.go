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
