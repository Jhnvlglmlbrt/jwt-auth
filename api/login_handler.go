package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/jwt-auth/usecase"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds usecase.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !usecase.VerifyPassword(creds.Username, creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenSigned, err := usecase.GenerateToken(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usecase.SetTokenCookie(w, tokenSigned)
}
