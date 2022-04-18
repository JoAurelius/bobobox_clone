package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}
