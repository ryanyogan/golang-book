package main

import "fmt"

func main() {
	fmt.Print("Please enter a temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	celcius := (input - 32) * 5 / 9

	fmt.Printf("%f in Celcius is: %f", input, celcius)
}
