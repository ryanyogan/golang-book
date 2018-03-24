package main

import "fmt"

func main() {
	var x string
	x = "Hello World"
	fmt.Println(x)

	r := 5
	fmt.Println(r)

	const z string = "Can't Change"
	fmt.Println(z)

	var (
		name = "Ryan"
		age  = "34"
	)

	fmt.Println("I am ", name, age)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
