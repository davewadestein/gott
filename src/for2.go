package main

import "fmt"

func main() {
	var num int = 10
	fmt.Print("Start the countdown at what number? ")
	fmt.Scanln(&num)

	for ; num > 0; num-- { // HL
		fmt.Println(num)
	} // HL
	fmt.Println("Blast off!")
}
