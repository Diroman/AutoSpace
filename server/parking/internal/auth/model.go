package auth

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
