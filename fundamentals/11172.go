package main

import "fmt"

func main() {
	var testCases, num1, num2 int
	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&num1, &num2)

			if err == nil {
				if num1 < num2 {
					fmt.Println("<")
				} else if num1 > num2 {
					fmt.Println(">")
				} else {
					fmt.Println("=")
				}
			} else {
				break
			}
		}
	}
}
