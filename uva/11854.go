package main

import "fmt"

func main() {
	var side1, side2, side3 int
	_, err := fmt.Scan(&side1, &side2, &side3)

	for err == nil {
		if side1 == 0 && side2 == 0 && side3 == 0 {
			break
		}

		if (side1 > side2 && side1 > side3 && side1*side1 == side2*side2+side3*side3) || (side2 > side1 && side2 > side3 && side2*side2 == side1*side1+side3*side3) || (side3 > side1 && side3 > side2 && side3*side3 == side1*side1+side2*side2) {
			println("right")
		} else {
			println("wrong")
		}

		_, err = fmt.Scan(&side1, &side2, &side3)
	}
}
