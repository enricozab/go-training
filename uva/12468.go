package main

import (
	"fmt"
)

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)

	for err == nil && a != -1 && b != -1 {
		if b >= a && b-a+1 <= 50 {
			fmt.Println(b - a)
		} else if b >= a && b-a+1 >= 50 {
			fmt.Println(99 - b + a + 1)
		} else if a >= b && a-b+1 <= 50 {
			fmt.Println(a - b)
		} else {
			fmt.Println(99 - a + b + 1)
		}

		_, err = fmt.Scan(&a, &b)
	}
}
