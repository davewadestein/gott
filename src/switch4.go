package main

// START OMIT
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	_, _, sec := time.Now().Clock() // get time, extract seconds, discard hours and minutes
	rand.Seed(int64(sec))           // use seconds to seed random number generator
	r := rand.Int()
	fmt.Printf("Random number (%d) is ", r)

	switch { // HL
	case r%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	} // HL
}

// END OMIT
