package main 
import "fmt"


// 1. Basic Function
// Simple function without parameters and return types

func greet() {
	fmt.Println("Hello from the greet function!")
}

func main (){
	greet() // calling the function
	add(5, 10) // calling the function with parameters

	result := multiply(4,9) // calling the function with return value
	fmt.Println("Multiplication Result:", result)

	quotient, remainder := divide(10, 3) // calling the function with multiple return values
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}

// 2. Function with Parameters
func add (a int, b int){
	fmt.Println("Sum:", a+b)
}

// 3. Function with Return Value
func multiply(a int, b int) int {
	return a * b
}

// 4. Function with Multiple Return Values
func divide(divided int, divisor int) (int, int){
	quotient := divided / divisor
	remainder := divided % divisor
	return quotient, remainder
}