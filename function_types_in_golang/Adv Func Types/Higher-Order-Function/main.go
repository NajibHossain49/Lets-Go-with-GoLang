package main

import "fmt"

// A higher-order function is a function that either takes one or more functions as arguments,
// or returns a function as its result, or both.
// Higher-order functions are a powerful feature of functional programming and can be used to create more abstract and reusable code.


// Function that takes another function as an argument
func applyOperation(z int, x int, operation func(a int, b int)) {
	operation(x, z)

}

// Function that returns another function
func multiplyOperation() func(a int, b int) {
	return multiply
}


func multiply(a int, b int) {
	z := a * b
	fmt.Println("Product:", z)
}

func add(a int, b int) { 
	z := a + b
	fmt.Println("Sum:", z)
}

func main() {
	applyOperation(5, 3, add)

	multiplyFunc := multiplyOperation()
	multiplyFunc(4, 6)
}


