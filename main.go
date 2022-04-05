package main

import (
	"fmt"
	"net/http"

	"rsym/go_http_sample/handler"
)

func main() {
	http.HandleFunc("/", handler.TopHandler)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/submit", handler.SubmitHandler)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe", err)
	}
}
