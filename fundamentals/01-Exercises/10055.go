package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	_, err := fmt.Scan(&num1, &num2)

	for err == nil {
		fmt.Println(int(math.Abs(num1 - num2)))
		_, err = fmt.Scan(&num1, &num2)
	}
}
