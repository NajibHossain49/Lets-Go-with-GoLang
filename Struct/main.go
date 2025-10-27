package main

import "fmt"

type Person struct {
    name string // member variables বা fields or property
    age  int    // member variables বা fields or property
    city string // member variables বা fields or property
}

func main() {
    // Struct-এর একটি instance তৈরি করা
    p1 := Person{"Najib", 25, "Dhaka"}

    // অথবা key-value ফরম্যাটে
    p2 := Person{name: "Rahim", age: 30, city: "Chittagong"}

    fmt.Println(p1)
    fmt.Println(p2)

    // নির্দিষ্ট field অ্যাক্সেস করা
    fmt.Println("Name:", p1.name)
    fmt.Println("Age:", p1.age)
}
