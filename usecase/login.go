package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("my_secret_key")
var expireTime = time.Now().Add(5 * time.Minute)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func VerifyPassword(username, password string) bool {
	expectedPassword, err := users[username]
	return err && expectedPassword == password
}

func GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}
