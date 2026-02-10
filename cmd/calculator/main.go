package main

import (
	"calculator/internal/arithmetic"
	"calculator/internal/storage"
	"fmt"
)

func main() {
	// Load existing history at startup
	historyList := storage.LoadHistory()
	fmt.Printf("Loaded %d previous calculations.\n", len(historyList))
	var num1, num2, result float64
	var operator string
	var err error

	for {
		// ... (your existing input logic) ...

		if operator == "exit" {
			storage.SaveHistory(historyList)
			break
		}

		result, err = arithmetic.Calculate(num1, num2, operator)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

			// 2. Add the successful calculation to your historyList
			calc := storage.Calculation{
				FirstNum:  num1,
				Operator:  operator,
				SecondNum: num2,
				Result:    result,
			}
			historyList = append(historyList, calc)
		}

		if err == nil {
			// Add new calculation to the loaded list
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
