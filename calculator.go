package main

import (
	"errors"
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

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	for {
		num1 := getValidNumber("Enter first number: ")

		var operator string
		for {
			fmt.Print("Enter operator (+, -, *, /, or 'exit' to quit): ")
			fmt.Scanln(&operator)

			// Check if the operator is one of the valid choices
			if operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "exit" {
				break // Valid input received, exit the operator loop
			}

			// If not valid, show error and the loop repeats the operator prompt
			fmt.Println("Error: Invalid operator! Please enter +, -, *, or /.")
		}

		// Check if they want to quit the entire program
		if operator == "exit" {
			fmt.Println("Closing program...")
			break // Breaks the outer main loop
		}

		num2 := getValidNumber("Enter second number: ")

	}

	// Now the switch only runs if the operator is guaranteed to be valid
		switch operator {
		case "+":
			fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)

		case "-":
			fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)

		case "*":
			fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)

		case "/":
			if num2 == 0 {
				fmt.Println("Undefined")
				// Handle division by zero
			} else {
				fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
				// Perform division
			}
		default:
			fmt.Println("test.")
		}
	}
}