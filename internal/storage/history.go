package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

	// Capture the error here to know if the file is corrupted
	err = json.Unmarshal(fileData, &history)
	if err != nil {
		fmt.Println("Warning: History file is corrupted. Starting with empty history.")
		return []Calculation{}
	}
	return history
}

func SaveHistory(history []Calculation) {
	jsonData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// This shows the exact path on your hard drive
	absPath, _ := filepath.Abs("data/history.json")
	fmt.Printf("Saving history to: %s\n", absPath)

	err = os.WriteFile("data/history.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
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
