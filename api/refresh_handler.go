package api

import (
	"net/http"
	"time"

	"github.com/Jhnvlglmlbrt/jwt-auth/usecase"
	"github.com/golang-jwt/jwt/v5"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := c.Value

	claims, tkn, err := usecase.ParseJWT(tokenStr)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Until(claims.ExpiresAt.Time) > 3*time.Minute {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expireTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expireTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(usecase.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usecase.SetTokenCookie(w, tokenString)
}
