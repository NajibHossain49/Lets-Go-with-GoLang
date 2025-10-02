package main

import "fmt"

var(
	a int = 10
	b int = 20
	c int = 30
)

func add(x int, y int)  {
 z:= x + y
 fmt.Println("Sum is:", z)
}

func main() {
q:=30
p:=40

	add(a,b)
	add(p,q)
	add(a,c)
	add(a,z)

}