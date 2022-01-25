package main

import (
	"fmt"
)

func main() {
	var input string
	var ctr = 1

	_, err := fmt.Scan(&input)

	for err == nil && input != "*" {
		var output string

		if input == "Hajj" {
			output = "Hajj-e-Akbar"
		} else if input == "Umrah" {
			output = "Hajj-e-Asghar"
		}

		fmt.Printf("Case %v: %v\n", ctr, output)
		ctr++

		_, err = fmt.Scan(&input)
	}
}
