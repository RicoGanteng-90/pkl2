package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Username string `json:"userName"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
