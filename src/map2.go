package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int) // To create an empty map, use make
	m["tall"] = 12            // Set key/value pairs
	m["grande"] = 16
	m["venti"] = 20
	fmt.Println(m)
	fmt.Println("len:", len(m)) // HL
	delete(m, "grande")         // HL
	fmt.Println(m)

	// Optional second return value indicates if key was in map.
	// Used to disambiguate between missing keys and zero values.
	_, present := m["grande"]
	fmt.Println("grande present?", present)
}

// END OMIT
