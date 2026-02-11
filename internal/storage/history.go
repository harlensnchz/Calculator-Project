package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type Calculation struct {
	FirstNum  float64 `json:"first_num"`
	Operator  string  `json:"operator"`
	SecondNum float64 `json:"second_num"`
	Result    float64 `json:"result"`
}

// This locks the file to your project folder for GitHub tracking
func getHistoryPath() string {
	return `C:\Users\Harlen\Documents\Calculator-Project\data\history.json`
}

func LoadHistory() []Calculation {
	path := getHistoryPath()
	var history []Calculation
	fileData, err := os.ReadFile(path)
	if err != nil {
		return []Calculation{}
	}
	json.Unmarshal(fileData, &history)
	return history
}

func SaveHistory(history []Calculation) {
	path := getHistoryPath()
	jsonData, _ := json.MarshalIndent(history, "", "  ")
	os.WriteFile(path, jsonData, 0644)
}

// NEW: Clear History Function
func ClearHistory() {
	path := getHistoryPath()
	// Overwrite the file with an empty JSON array
	err := os.WriteFile(path, []byte("[]"), 0644)
	if err != nil {
		fmt.Println("Error clearing history:", err)
	} else {
		fmt.Println("History cleared successfully!")
	}
}

func ShowHistory(history []Calculation) {
	fmt.Println("\n--- ALL SAVED CALCULATIONS ---")
	if len(history) == 0 {
		fmt.Println("No history found.")
		return
	}
	for i, entry := range history {
		fmt.Printf("[%d] %.2f %s %.2f = %.2f\n", i+1, entry.FirstNum, entry.Operator, entry.SecondNum, entry.Result)
	}
	fmt.Println("------------------------------")
}
