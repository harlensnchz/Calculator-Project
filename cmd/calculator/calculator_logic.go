package main // Important: This must stay 'package main' to talk to main.go

import (
	"calculator/internal/arithmetic"
	"calculator/internal/storage"
	"fmt"
)

// Move your runCalculator and getValidNumber functions here
func runCalculator(historyList *[]storage.Calculation) {
	for {
		num1 := getValidNumber("Enter first number: ")

		var operator string
		for {
			fmt.Print("Enter operator (+, -, *, /, %, or 'back' to menu): ")
			fmt.Scanln(&operator)

			if operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "%" || operator == "back" {
				break
			}
			fmt.Println("Error: Invalid operator!")
		}

		if operator == "back" {
			return
		}

		num2 := getValidNumber("Enter second number: ")
		result, err := arithmetic.Calculate(num1, num2, operator)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

			calc := storage.Calculation{FirstNum: num1, Operator: operator, SecondNum: num2, Result: result}
			*historyList = append(*historyList, calc)
		}
	}
}
