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
		num1 := getValidNumber("Enter first number: ")

		var operator string
		fmt.Print("Enter operator (+, -, *, /, or 'exit' to quit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			break
		}

		// --- NEW VALIDATION LOGIC ---
		// Check if the operator is valid BEFORE asking for the second number
		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Error: Invalid operator! Please enter +, -, *, or /.")
			continue // This skips the rest of the loop and restarts from the top
		}

		num2 := getValidNumber("Enter second number: ")

		// Now the switch only runs if the operator is guaranteed to be valid
		switch operator {
		case "+":
			fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)

		case "-":
			fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)

		case "*":
			fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)

		case "/":
			fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)

		}
	}
}
