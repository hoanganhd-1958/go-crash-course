package main

import "fmt"

func main() {
	fmt.Println("Hello")
	variable()
}
func variable() {
	var age int
	age = 21
	fmt.Println("my age is", age)
	name, age := "naveen", 90

	fmt.Println("my name is", name, "age is", age)
}
