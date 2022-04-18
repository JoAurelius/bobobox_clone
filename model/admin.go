package model

type Admin struct {
	AdminID       int    `json:"admin_id,omitempty"`
	AdminName     string `json:"admin_name,omitempty"`
	AdminPassword string `json:"admin_password,omitempty"`
}

type AdminResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Admin  `json:"data"`
}
type AdminsResponses struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Admin `json:"data"`
}
