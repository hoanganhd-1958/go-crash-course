package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Hello, you've requested: %s\n\n", request.URL.Path)
		fmt.Printf("Hello, you've requested: %s\n\n", request.Method)
	})

	http.ListenAndServe(":8080", nil)
}
