package model

type Admin struct {
	AdminID       string `json:"admin_id"`
	AdminPassword string `json:"admin_password"`
}

type AdminResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Admin  `json:"data"`
}
