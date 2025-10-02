package main

import "fmt"

func main() {
// Explicit type
var age int = 25

// Type inference (compiler infers the type)
var name = "Alice"

// Short declaration (only inside functions)
score := 95

// Multiple variables
var x, y, z int = 1, 2, 3
var a, b = 10, "GoLang"
	
fmt.Println(a)
fmt.Println(b)
fmt.Println(x, y, z)
fmt.Println(age)
fmt.Println(name)
fmt.Println(score)
}


