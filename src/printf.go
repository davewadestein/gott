package main

// START OMIT
import (
	"fmt"  // we import packages in order to use them
	"math" // we'll use the math package too
)

func main() {
	var name, age = "Ranveer Singh", 33
	fmt.Printf("%s\n", "Hello, world!")
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Print("Enter an integer: ")

	num := 128881
	fmt.Printf("The square root of %d is %f\n", num, math.Sqrt(float64(num)))
}

// END OMIT
