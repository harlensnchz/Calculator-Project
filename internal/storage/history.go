package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// Calculation must be defined here and capitalized to be exported
type Calculation struct {
	FirstNum  float64 `json:"first_num"`
	Operator  string  `json:"operator"`
	SecondNum float64 `json:"second_num"`
	Result    float64 `json:"result"`
}

func LoadHistory() []Calculation {
	os.MkdirAll("data", 0755)
	var history []Calculation

	fileData, err := os.ReadFile("data/history.json")
	if err != nil {
		return []Calculation{}
	}

	json.Unmarshal(fileData, &history)
	return history
}

// Ensure this starts with a Capital 'S'
func SaveHistory(history []Calculation) {
	jsonData, _ := json.MarshalIndent(history, "", "  ")
	err := os.WriteFile("data/history.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving history:", err)
	}
}

func ShowHistory(history []Calculation) {
	fmt.Println("\n--- ALL SAVED CALCULATIONS ---")
	if len(history) == 0 {
		fmt.Println("No history found in data/history.json.")
		return
	}

	// Iterating through the history slice you loaded at startup
	for i, entry := range history {
		fmt.Printf("[%d] %.2f %s %.2f = %.2f\n", i+1, entry.FirstNum, entry.Operator, entry.SecondNum, entry.Result)
	}
	fmt.Println("------------------------------")
}
