Go CLI Calculator

A terminal-based calculator built with Go (Golang) that prioritizes data integrity and a smooth user experience. This project demonstrates modular function design, input buffer management, and strict control flow validation.

Key Features

Modular Validation: Includes a custom getValidNumber function that uses a for loop to ensure the program only proceeds when numeric data is provided.

Buffer Clearing: Prevents "infinite skip" bugs by manually clearing the input buffer (fmt.Scanln(&discard)) whenever non-numeric characters are detected.

Early Exit & Gatekeeping: Validates the mathematical operator immediately after entry. If an invalid operator is typed, the program uses continue to reset the cycle without wasting time asking for a second number.

Persistent Session: Wrapped in a global for loop with a dedicated exit command to allow for back-to-back calculations.

Formatted Output: Uses %.2f formatting to provide clean, readable results rounded to two decimal places.

Technical Implementation

Input Validation Logic
The program uses the err return value from fmt.Scanln to detect data type mismatches:

Go
_, err := fmt.Scanln(&num)
if err == nil {
    return num
}

If err is not nil, the program identifies that a character was entered instead of a float64, notifies the user, and clears the terminal buffer.

Control Flow

Instead of a simple linear script, this project uses Conditional Gatekeeping. The operator input acts as a checkpoint:

If exit: The program triggers a break.

If Invalid: The program triggers a continue (skipping the num2 prompt).

If Valid: The program proceeds to capture the second operand and execute the switch logic.

How to Run

Ensure you have Go installed on your system.

Clone this repository.

Run the following command in your terminal:

Bash
go run calculator.go

Developer: Harlenken Sanchez

Academic Status: Graduating BSIT Student (4th Year)
