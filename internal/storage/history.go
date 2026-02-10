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
	fmt.Println("\n--- Calculation History ---")
	if len(history) == 0 {
		fmt.Println("No history found.")
		return
	}
	for i, c := range history {
		fmt.Printf("[%d] %.2f %s %.2f = %.2f\n", i+1, c.FirstNum, c.Operator, c.SecondNum, c.Result)
	}
	fmt.Println("---------------------------")
}
