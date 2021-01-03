package main

import (
	"fmt"
	"net/http"
)

// https://golangbot.com/webassembly-using-go/

func main() {
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../../assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
