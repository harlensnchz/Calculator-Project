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

func SaveHistory(history []Calculation) {
	jsonData, _ := json.MarshalIndent(history, "", "  ")

	// Note: Path updated to your new 'data' folder
	err := os.WriteFile("data/history.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving history:", err)
	}
}
