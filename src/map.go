package main

import "fmt"

// START OMIT
func main() {
	continents := map[string]string{
		"United States": "North America",
		"India":         "Asia",
		"Namibia":       "Africa",
		"Germany":       "Europe",
		"Mexico":        "North America",
		"Ireland":       "Europe",
		"Iceland":       "Europe",
	}
	fmt.Printf("%#v\n", continents)
	c := "Iceland"
	fmt.Println(c, "is in", continents[c])
	c = "Freedonia"
	fmt.Println(c, "is in", continents[c])
}

// END OMIT
