package main

import "fmt"

// START OMIT
func main() {
    f()
    fmt.Println("Back from f()")
}

func f() {
    defer fmt.Println("Defer in f")
    fmt.Println("Calling g()")
    g(0)
    fmt.Println("Back from g()")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic("buffer underrun") // HL
    }
    defer fmt.Println("Defer in g", i) // HL
    fmt.Println("Printing in g", i)
    g(i + 1)
}
//END OMIT
