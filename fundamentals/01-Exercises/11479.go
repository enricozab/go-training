package main

import "fmt"

func main() {
	var testCases, side1, side2, side3 int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&side1, &side2, &side3)

			if err == nil {
				if (side1 == 0 || side2 == 0 || side3 == 0) || (side1+side2 < side3 || side1+side3 < side2 || side2+side3 < side1) {
					fmt.Printf("Case %v: %v\n", ctr+1, "Invalid")
				} else if side1 == side2 && side2 == side3 {
					fmt.Printf("Case %v: %v\n", ctr+1, "Equilateral")
				} else if side1 == side2 || side1 == side3 || side2 == side3 {
					fmt.Printf("Case %v: %v\n", ctr+1, "Isosceles")
				} else {
					fmt.Printf("Case %v: %v\n", ctr+1, "Scalene")
				}
			} else {
				break
			}
		}
	}
}
