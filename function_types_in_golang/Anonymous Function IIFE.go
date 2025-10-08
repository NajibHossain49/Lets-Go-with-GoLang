// এমন একটি ফাংশন যার কোনো নাম থাকে না।
package main

import "fmt"

func main() {
	// Anonymous function assign করা হলো একটি ভেরিয়েবলে
	add := func(a int, b int) int {
		return a + b
	}

	// এখন ভেরিয়েবল দিয়ে ফাংশনকে কল করা হচ্ছে
	result := add(10, 20)
	fmt.Println("Result:", result)

	// IIFE (Immediately Invoked Function Expression)
	func(a int, b int) {
		sum := a + b
		fmt.Println("Immediate Result:", sum)
	}(15, 25)
}



