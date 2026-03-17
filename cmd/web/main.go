package main

import (
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", http.HandlerFunc(home))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	
	log.Printf("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}