package arithmetic

import "errors"

// Calculate performs the arithmetic operations.
// Note: We capitalize "Calculate" so it can be used in main.go.
func Calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	case "%":
		if num2 == 0 {
			return 0, errors.New("modulus by zero")
		}
		return float64(int(num1) % int(num2)), nil
	default:
		return 0, errors.New("invalid operator")
	}
}
