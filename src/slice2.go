package main

import "fmt"

// START OMIT
func main() {
	s := make([]string, 3) // slice of strings of size 3
	fmt.Println("initial slice:", s)
	s[0], s[1], s[2] = "zero", "one", "two" // set just like an array
	fmt.Println("now:", s)
	fmt.Println("get:", s[2]) // get like an array
	fmt.Println("len of slice:", len(s))
	s = append(s, "three")        // HL
	s = append(s, "four", "five") // HL
	fmt.Println("after append:", s)
	copy_of_s := make([]string, len(s)) // HL
	copy(copy_of_s, s)                  // HL
	fmt.Println("copy:", copy_of_s)
}

// END OMIT
