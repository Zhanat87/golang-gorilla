package server

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


