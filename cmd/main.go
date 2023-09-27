package main

import (
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/jwt-auth/api"
)

func main() {
	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/welcome", api.Welcome)
	http.HandleFunc("/refresh", api.Refresh)
	http.HandleFunc("/logout", api.Logout)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
