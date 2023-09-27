package usecase

import "net/http"

func SetTokenCookie(w http.ResponseWriter, tokenStr string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expireTime,
	})
}
