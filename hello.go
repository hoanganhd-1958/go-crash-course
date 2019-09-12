package main

import "fmt"

func main() {
	fmt.Println("Hello")
	variable()
	fmt.Println(calculateBill(1000, 21))
}

func variable() {
	var age int
	age = 21
	fmt.Println("my age is", age)
	name, age := "naveen", 90

	fmt.Println("my name is", name, "age is", age)
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	area, _, h := rectProps2(10.8, 5.6, 2.3)
	fmt.Printf("Area %f Perimeter %f", area, h)
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64, h float64) (float64, float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter, h
}
