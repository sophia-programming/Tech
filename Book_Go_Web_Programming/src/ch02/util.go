package main

import (
	"net/http"

	"main.go/data"
)

func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = error.New("Invalid session")
		}
	}
	return
}
