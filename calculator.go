package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter Your First Number:")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("Enter Operator (+, -, *, /):")
	var operator string
	fmt.Scanln(&operator)

	fmt.Println("Enter Your Second Number:")
	var num2 float64
	fmt.Scanln(&num2)

	// CALLING A FUNCTION AND SAVING RESULT!
	cal := calculate(num1, operator, num2)
	fmt.Println("Result:", cal)

}

// CALCULATE FUNCTION!
func calculate(num1 float64, operator string, num2 float64) float64 {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 != 0 {
			return num1 / num2
		} else {
			fmt.Println("Error: Cannot divide by zero.")
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid operator.")
		os.Exit(1)
		return 0
	}

	return 0
}
