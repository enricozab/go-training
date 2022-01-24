package main

import (
	"fmt"
)

func main() {
	var testCases, l, w, h int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 1; ctr <= testCases; ctr++ {
			_, err = fmt.Scan(&l, &w, &h)

			if err == nil {
				var condition = "bad"
				if l <= 20 && w <= 20 && h <= 20 {
					condition = "good"
				}

				fmt.Printf("Case %v: %v\n", ctr, condition)
			} else {
				break
			}
		}
	}
}
