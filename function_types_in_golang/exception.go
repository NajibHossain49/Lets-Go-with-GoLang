package main

import "fmt"

func main() {
    fmt.Println("Hello, World! from main function")

	add(5, 10) // Calling the add function first for global scope function calling

	add:= func(a int, b int) {
		fmt.Println("Subtraction:", a - b)
	}
	
	add(10, 5) // Calling the add function second for local scope function calling
}

func init () {
	fmt.Println("Init function called")
}

// func add(a int, b int) {
// 	fmt.Println("Addition:", a + b)
// }