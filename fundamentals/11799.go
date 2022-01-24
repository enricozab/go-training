package main

import (
	"fmt"
)

func main() {
	var testCases, inputs, speed, max int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 1; ctr <= testCases; ctr++ {
			max = 0
			_, err = fmt.Scan(&inputs)

			if err == nil {
				for ctr := 0; ctr < inputs; ctr++ {
					_, err = fmt.Scan(&speed)

					if err == nil && speed > max {
						max = speed
					}
				}

				fmt.Printf("Case %v: %v\n", ctr, max)
			} else {
				break
			}
		}
	}
}
