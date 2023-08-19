package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculator 1.0")
	fmt.Println("--------------")
	fmt.Println("Which operation would you like to perform? (add, sub, mul, div)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "sub":
		fmt.Println(num1 - num2)
	case "mul":
		fmt.Println(num1 * num2)
	case "div":
		fmt.Println(num1 / num2)
	}
}
