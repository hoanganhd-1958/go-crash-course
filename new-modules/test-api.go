package main

import (
	"io/ioutil"
	"net/http"
)
import "fmt"

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		// handle error
	}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}