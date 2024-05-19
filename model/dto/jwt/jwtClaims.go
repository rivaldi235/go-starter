package jwtClaims

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Roles    string `json:"role"`
}
