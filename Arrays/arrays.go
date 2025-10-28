package main

import "fmt"

var (
	arr1 = [5]int{1, 2, 3, 4, 5} 
)

func main() {

	arr := [5]int{10, 20, 30, 40, 50}
	

    arr[1] = 100
	fmt.Println("Array initialized:", arr)
	fmt.Println("Array initialized:", arr1[4])
	fmt.Println("Array initialized:", arr1)
}

