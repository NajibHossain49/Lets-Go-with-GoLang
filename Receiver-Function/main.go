package main 

import "fmt"


type Person struct {
	Name string
	Age  int
}

func PrintPersonInfo(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// Receiver function to print details of Person
func (p Person)PrintDetails(height int){
	fmt.Println("Person Details - Name:", p.Name, ", Age:", p.Age, ", Height:", height)
}

func main(){
	var p Person

	p = Person{
		Name: "Alice",
		Age:  30,
	}

	p.PrintDetails(10)
	PrintPersonInfo(p)
}