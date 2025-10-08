// Init Function

package main

import "fmt"

var x int = 10

func init() {
	fmt.Println("Initialization, I'm the first function to be executed")
	// init() // This will cause a compile-time error: init function cannot be called explicitly

	fmt.Println("Value of x:", x)
	x = 20

}

func main() {
	fmt.Println("Main function")
	fmt.Println("Value of x in main:", x)
}
