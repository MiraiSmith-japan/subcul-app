package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pongHandler)
	http.ListenAndServe(":1991", nil);
	fmt.Print("start server")
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
	w.WriteHeader(200)
}