package main

import (
	"fmt"
	"strconv"
)

var num int = 451

func main() {
	switch strconv.Itoa(num) {
	case "451", "452":
		fmt.Println("Worked!")
	default:
		fmt.Println("Failed!")
	}
}
