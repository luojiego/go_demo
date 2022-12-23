package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<H1>Hello Roger!</H1>"))
	})

	http.ListenAndServe(":9090", nil)
}
