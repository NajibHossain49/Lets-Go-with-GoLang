package main

import ("fmt"
"example.com/mathLib"
)

var a = 10 // package-level variable

func main() {
    b := 20 // function-level variable
    fmt.Println("a:", a) // বাইরের ভ্যারিয়েবল অ্যাক্সেস করা যাবে
    fmt.Println("b:", b)
    // add (5, 10)

    mathLib.Add(15, 25) // Custom package function call

	
	fmt.Println("Using Custom package variable:", mathLib.Name)
	
}