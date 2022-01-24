package main

import "fmt"

func main() {
	var testCases, salary1, salary2, salary3 int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&salary1, &salary2, &salary3)

			if err == nil {
				if (salary1 > salary2 && salary1 < salary3) || (salary1 < salary2 && salary1 > salary3) {
					fmt.Printf("Case %v: %v\n", ctr+1, salary1)
				} else if (salary2 > salary1 && salary2 < salary3) || (salary2 < salary1 && salary2 > salary3) {
					fmt.Printf("Case %v: %v\n", ctr+1, salary2)
				} else {
					fmt.Printf("Case %v: %v\n", ctr+1, salary3)
				}
			} else {
				break
			}
		}
	}
}
