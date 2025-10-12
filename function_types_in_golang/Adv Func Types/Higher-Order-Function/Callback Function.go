package main

import "fmt"

// Higher-Order Function
// এটি একটি function কে parameter হিসেবে নেয় (callback function)
func processNumbers(a int, b int, callback func(int, int) int) {
	result := callback(a, b) // এখানে callback function কে call করা হচ্ছে
	fmt.Println("Result:", result)
}

// Callback functions
func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func main() {
	// processNumbers হচ্ছে Higher-Order Function
	// add এবং multiply হচ্ছে Callback Function

	processNumbers(10, 5, add)      // add function কে callback হিসেবে পাঠানো হয়েছে
	processNumbers(10, 5, multiply) // multiply function কে callback হিসেবে পাঠানো হয়েছে
}
