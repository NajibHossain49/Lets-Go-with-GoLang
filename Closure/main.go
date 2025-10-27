package main

import "fmt"

func counter() func() int {
    count := 0 // outer variable

	show := func() int {
        count++
        return count
    }
    return show
}

func main() {
    next := counter()
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3

    another := counter()
    fmt.Println(another()) // 1 (new independent closure)
}
