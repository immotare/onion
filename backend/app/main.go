package main

import (
	"log"
	"net/http"
)

func OkResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("response sucess"))
	log.Println("response returned")
}

func main() {
	http.HandleFunc("/", OkResponseHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println("failed starting server")
	}
}
