package main

import "fmt"

// this function cannot change the variable passed in to it
func increment(num int) { 
    num++
}

func main() {
    val := 1
    increment(val) 
    fmt.Println("Value after incrementing:", val)
}
