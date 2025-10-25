package main

import "fmt"

func add(a, b int) int {
    c := a + b
    return c
}

func main() {
    result := add(5, 6) // store the returned value in a variable
    fmt.Print("Result:", result)
}