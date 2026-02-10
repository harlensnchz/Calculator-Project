package main

import (
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
	historyList := storage.LoadHistory() //

	for {
		fmt.Println("\n\n=== CALCULATOR MENU ===")
		fmt.Println("\n1. Use Calculator")
		fmt.Println("2. View History")
		fmt.Println("3. Exit")
		fmt.Print("\nSelect an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			runCalculator(&historyList) // Now this function exists!
		case "2":
			storage.ShowHistory(historyList)
			fmt.Print("Press Enter to return to menu...")
			fmt.Scanln()
		case "3":
			storage.SaveHistory(historyList)
			fmt.Println("Thank you for using the calculator!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
