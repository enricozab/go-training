package main

import (
	"fmt"
)

func main() {
	var testCases, farmers, size, animals, environment int

	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&farmers)

			if err == nil {
				sum := 0

				for ctr2 := 0; ctr2 < farmers; ctr2++ {
					_, err = fmt.Scan(&size, &animals, &environment)

					if err == nil {
						sum += size * environment
					} else {
						break
					}
				}

				println(sum)
			} else {
				break
			}
		}
	}
}
