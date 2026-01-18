package main

import (
	"calculator/operations"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go [add | subtract | multiply | divide] [number1] [number2]")
		return
	}


	operation := os.Args[1]
	num1, _ := strconv.ParseFloat(os.Args[2], 64)
	num2, _ := strconv.ParseFloat(os.Args[3], 64)

	switch operation {
		case "add":
			fmt.Printf("Result: %.2f\n", operations.Add(num1, num2))
		case "subtract":
			fmt.Printf("Result: %.2f\n", operations.Subtract(num1, num2))
		case "multiply":
			fmt.Printf("Result: %.2f\n", operations.Multiply(num1, num2))
		case "divide":
			result, err := operations.Divide(num1, num2)
			if err != nil {
				fmt.Printf("Result %f\n", result)
			} else {
				fmt.Printf("Error: %v\n", result)
			}
		default:
			fmt.Println("Invalid operation")
			
	}
}