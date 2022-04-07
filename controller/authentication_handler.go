package controller

import (
	"fmt"
	"net/http"
	"time"

	"bobobox_clone/model"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("asdlk12AS12dw29")
var tokenName = "bobotoken"

func generateToken(w http.ResponseWriter, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(30 * time.Minute)

	claims := &model.Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		SendGeneralResponse(w, 401, "tokensignederror")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: false,
		Path:     "/",
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    tokenName,
		Value:   "",
		Expires: time.Now(),
		// Secure:   false,
		// HttpOnly: false,
		Path: "/",
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidate := validateUserToken(r, accessType)
		if !isValidate {
			SendGeneralResponse(w, 401, "not autorized")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, id, email, userType := validateTokenFromCookies(r)
	fmt.Print(id, email, userType, isAccessTokenValid)

	if isAccessTokenValid {
		isUservalid := userType == accessType
		if isUservalid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, int, string, int) {
	cookie, err := r.Cookie(tokenName)
	if err == nil {
		accessToken := cookie.Value
		accessClaim := &model.Claims{}
		parseToken, err := jwt.ParseWithClaims(accessToken, accessClaim, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parseToken.Valid {
			return true, accessClaim.ID, accessClaim.Name, accessClaim.UserType
		}
	}
	return false, -1, "", -1
}
