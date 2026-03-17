package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){
	port := flag.String("port", ":4000", "HTTP port")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", http.HandlerFunc(home))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	
	log.Printf("Starting server on port %s", *port)

	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}