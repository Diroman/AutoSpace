package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"
)

const jwtKey = "secret"

func CreateNewToken(id int) (string, error) {
	claims := &Claims{
		ID: strconv.Itoa(id),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5*time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(token string) (int, bool) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, false
		}
		return http.StatusBadRequest, false
	}
	if !tkn.Valid {
		return http.StatusUnauthorized, false
	}

	id, _ := strconv.Atoi(claims.ID)
	return id, true
}