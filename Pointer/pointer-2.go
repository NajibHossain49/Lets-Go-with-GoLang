package main

import "fmt"

// Defining a struct person type
type Person struct {
	name   string
	age    int
	height int
}



func printArray(number *[3]int){ // pointer to an array of 3 integers

	 fmt.Println(number)
}

func main(){
	arr := [3]int{10, 20, 30}
   printArray(&arr) // passing the address of the array

//  initializing struct
	p1 := Person{name: "Alice", age: 30, height: 165}
	// p1 := "Hello, Go!"

	fmt.Println("Just normal struct's value printing:", p1)

	Address := &p1 // pointer to struct to find address of struct

    fmt.Println("Address of p1:", Address)
	
	*Address = p1

	fmt.Println("Value at address of p1:", *Address)
	

}