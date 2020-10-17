package auth

import (
	"github.com/dgrijalva/jwt-go"
	"math"
	"net/http"
)

const jwtKey = "super(secret_token)key"

func CreateNewToken(id int) (string, error) {
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: math.MaxInt64,
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

	return 0, true
}