package main

import "fmt"

func main() {
	var testCases, range1, range2 int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&range1, &range2)

			if err == nil && range1 <= range2 {
				var sum int

				for range1 <= range2 {
					if range1%2 != 0 {
						sum += range1
					}

					range1++
				}

				fmt.Printf("Case %v: %v\n", ctr+1, sum)
			} else {
				fmt.Printf("Case %v: Invalid input\n", ctr+1)
			}
		}
	}
}
