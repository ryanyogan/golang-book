package main

import "fmt"

func main() {
	fmt.Print("Please enter a distance in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048
	fmt.Printf("%f in meters is: %f", input, meters)
}
