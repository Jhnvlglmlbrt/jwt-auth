package usecase

import (
	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(tokenStr string) (*Claims, *jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	return claims, token, nil
}
