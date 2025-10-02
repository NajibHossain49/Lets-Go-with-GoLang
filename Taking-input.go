package main

import "fmt"

// Function to take a name as input and return it
func getName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}


// Function to take two numbers as input and return them
func getNumbers() (int, int) {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

// Function to calculate the sum of two numbers
func calculateSum(num1, num2 int) int {
	return num1 + num2
}

// Function to display the result
func showResult(name string, sum int) {
	fmt.Println("Hello", name)
	fmt.Println("The sum of the two numbers is:", sum)
}

func main() {
	name := getName()
	num1, num2 := getNumbers()
	sum := calculateSum(num1, num2)
	showResult(name, sum)
}

	