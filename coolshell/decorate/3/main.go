package main

import (
	"fmt"
	"log"
	"net/http"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("----> WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n",
		r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello world! "+r.URL.Path)
}

func main() {
	http.HandleFunc("/v1/hello", WithServerHeader(hello))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
