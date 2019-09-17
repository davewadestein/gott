package main

import "fmt"

const (
	Active, Running  = iota, iota + 100 // 0, 100
	Passive, Stopped                    // 1, 101

	// Reset // won't work because two expressions per line are needed

	// But, you can reset the last expression
	_ = iota
	Reset

	AnyOther = 10 // You can use any other expression even without iota
)

func main() {
	fmt.Println(Active, Running, Passive, Stopped, Reset, AnyOther)
}
