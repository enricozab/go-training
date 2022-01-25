package main

import (
	"fmt"
)

func main() {
	var testCases, students, grade int

	_, err := fmt.Scan(&testCases)

	if err == nil {
		for ctr := 0; ctr < testCases; ctr++ {
			_, err = fmt.Scan(&students)

			if err == nil {
				sum := 0
				count := 0
				var average int
				var percentage float64
				var grades []int

				for ctr2 := 0; ctr2 < students; ctr2++ {
					_, err = fmt.Scan(&grade)

					if err == nil {
						grades = append(grades, grade)
						sum += grade
					} else {
						break
					}
				}

				average = sum / students

				for _, val := range grades {
					if val > average {
						count++
					}
				}

				percentage = float64(count) / float64(students) * 100

				fmt.Printf("%.3f%%\n", percentage)
			} else {
				break
			}
		}
	}
}
