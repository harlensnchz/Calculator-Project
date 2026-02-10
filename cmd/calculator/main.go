package main

import (
	"fmt"

	"github.com/harlensnchz/calculator-project/internal/arithmetic"
	"github.com/harlensnchz/calculator-project/internal/storage"
)

var historyList []storage.Calculation

func getValidNumber(prompt string) float64 {
	for {
		var num float64
		fmt.Print(prompt)
		_, err := fmt.Scanln(&num)
		if err == nil {
			return num
		}
		fmt.Println("Error: Please enter a valid number.")
		var discard string
		fmt.Scanln(&discard)
	}
}

func main() {
	for {
		num1 := getValidNumber("Enter first number: ")

		var operator string
		fmt.Print("Enter operator (+, -, *, /, %, or 'exit'): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			storage.SaveHistory(historyList) // Save before quitting
			break
		}

		num2 := getValidNumber("Enter second number: ")

		// Call arithmetic package
		result, err := arithmetic.Calculate(num1, num2, operator)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f\n", result)

			// Track history using the storage struct
			calc := storage.Calculation{FirstNum: num1, Operator: operator, SecondNum: num2, Result: result}
			historyList = append(historyList, calc)
		}
	}
}
