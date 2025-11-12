package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number")
	fmt.Scanln(&num1)

	fmt.Println("enter the operator")
	fmt.Scanln(&operator)

	fmt.Println("enter the second number")
	fmt.Scanln(&num2)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		}
	}
	fmt.Println("final result is:", result)
}
