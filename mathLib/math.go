package mathLib

import "fmt"

var Name string = "Math Library"

func Add(num1 int, num2 int)  {
	z := num1 + num2
	fmt.Println("From Custom Package, Sum:", z)
}