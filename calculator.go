package main

import (
	"fmt"
)

// getValidNumber is a custom function that handles input validation
func getValidNumber(prompt string) float64 {
	for {
		var num float64
		fmt.Print(prompt)
		_, err := fmt.Scanln(&num)

		if err == nil {
			return num // Successfully got a number, return it to the caller
		}

		// Logic to handle invalid characters
		fmt.Println("Error: Please enter a valid number, not a character.")

		// Clear the buffer to prevent infinite loops
		var discard string
		fmt.Scanln(&discard)
	}
}

func main() {
	for {
		fmt.Println("\n--- Personal Calculator Project ---")

		// Use the custom function for both numbers
		num1 := getValidNumber("Enter first number: ")

		var operator string
		fmt.Print("Enter operator (+, -, *, /, or 'exit' to quit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("Closing program...")
			break
		}

		num2 := getValidNumber("Enter second number: ")

		// Processing logic
		switch operator {
		case "+":
			fmt.Printf("Result: %.4f + %.4f = %.4f\n", num1, num2, num1+num2)
		case "-":
			fmt.Printf("Result: %.4f - %.4f = %.4f\n", num1, num2, num1-num2)
		case "*":
			fmt.Printf("Result: %.4f * %.4f = %.4f\n", num1, num2, num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("Result: %.4f / %.4f = %.4f\n", num1, num2, num1/num2)
			} else {
				fmt.Println("Error: Undefined (Division by zero).")
			}
		default:
			fmt.Println("Invalid operator! Use +, -, *, or /.")
		}
	}
}
