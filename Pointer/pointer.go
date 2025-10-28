package main

import "fmt"

func main(){
	var a int = 42

	p := &a

	fmt.Println("Value of a:", a)

	fmt.Println("Address of a:", p)
	
	*p = 30

     fmt.Println("Value of a:", a)

	fmt.Println("Address of a:", p)
	fmt.Println("Value at address p:", *p)
}
