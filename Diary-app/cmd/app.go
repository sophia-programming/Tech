package main

import (
	"net/http"

	"Tech/Diary-app/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/upload", handler.UploadHandler)
	http.HandleFunc("/show", handler.ShowHandler)
	http.ListenAndServe(":8888", nil)
}
