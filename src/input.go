package main

import "fmt"

func main() {
    var input string
    fmt.Print("Enter text: ")
    fmt.Scanln(&input)
    fmt.Println("I think you just entered...", input)
}
