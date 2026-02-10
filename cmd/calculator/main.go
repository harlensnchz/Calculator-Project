package main

import (
	"calculator/internal/arithmetic"
	"calculator/internal/storage"
	"fmt"
)

// getValidNumber handles input validation to ensure the program doesn't crash on characters
func getValidNumber(prompt string) float64 {
	for {
		var num float64
		fmt.Print(prompt)
		_, err := fmt.Scanln(&num)

		if err == nil {
			return num
		}

		fmt.Println("Error: Please enter a valid number, not a character.")
		var discard string
		fmt.Scanln(&discard)
	}
}

func main() {
	// 1. Load existing history at startup
	historyList := storage.LoadHistory()
	fmt.Printf("Loaded %d previous calculations.\n", len(historyList))

	var num1, num2, result float64
	var operator string
	var err error

	for {
		// 2. PAUSE FOR INPUT: This stops the infinite loop
		num1 = getValidNumber("Enter first number: ")

		for {
			fmt.Print("Enter operator (+, -, *, /, %, or 'exit'): ")
			fmt.Scanln(&operator)

			if operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "%" || operator == "exit" {
				break
			}
			fmt.Println("Error: Invalid operator!")
		}

		if operator == "exit" {
			// 3. SAVE HISTORY: Persists your data to data/history.json
			storage.SaveHistory(historyList)
			fmt.Println("Closing program...")
			break
		}

		num2 = getValidNumber("Enter second number: ")

		// 4. CALCULATE: Calls your modular arithmetic package
		result, err = arithmetic.Calculate(num1, num2, operator)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

			// 5. APPEND: Adds the result to your tracking list
			calc := storage.Calculation{
				FirstNum:  num1,
				Operator:  operator,
				SecondNum: num2,
				Result:    result,
			}
			historyList = append(historyList, calc)
		}
	}
}
