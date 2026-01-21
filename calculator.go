package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("--- Personal Calculator Project ---")

	// 1. Capture user input for the first number
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	// 2. Capture the mathematical operator
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// 3. Capture user input for the second number
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// 4. Logic to perform the calculation
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Undefined.")
		}
	default:
		fmt.Println("Invalid operator! Please use +, -, *, or /.")
	}
}
