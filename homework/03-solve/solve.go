package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	// Checks if URL path is correct
	if r.URL.Path != "/solve" {
		log.Fatalln("Incorrect URL.")
	}

	if v, ok := r.Form["coef"]; ok {
		// Splits the string of numbers using `,` and store in a slice
		splittedStringNumbers := strings.Split(v[0], ",")

		var numbers []int
		allIntegers := true

		// Checks if all numbers in the string are integer and store the converted integer to a new slice
		for _, num := range splittedStringNumbers {
			isInt, intValue := IsInteger(num)

			numbers = append(numbers, intValue)
			if !isInt {
				allIntegers = false
			}
		}

		// Checks if there are 12 numbers and all are integers
		if len(splittedStringNumbers) != 12 || !allIntegers {
			fmt.Fprintf(w, "%s\n", "Invalid parameter. Parameter `coef` must contain 12 integers.")
		} else {
			// Prints the system of equations
			fmt.Fprintln(w, "System:")
			for index := 0; index < len(numbers); index += 4 {
				fmt.Fprintf(w, "%dx + %dy + %dz = %d\n", numbers[index], numbers[index+1], numbers[index+2], numbers[index+3])
			}
			fmt.Fprintln(w, "")

			// Computes for determinant, determinant-x, determinant-y, determinant-z
			var d, dx, dy, dz float64
			d = float64((numbers[0] * (numbers[5]*numbers[10] - numbers[6]*numbers[9])) - (numbers[4] * (numbers[1]*numbers[10] - numbers[2]*numbers[9])) + (numbers[8] * (numbers[1]*numbers[6] - numbers[2]*numbers[5])))
			dx = float64((numbers[3] * (numbers[5]*numbers[10] - numbers[6]*numbers[9])) - (numbers[7] * (numbers[1]*numbers[10] - numbers[2]*numbers[9])) + (numbers[11] * (numbers[1]*numbers[6] - numbers[2]*numbers[5])))
			dy = float64((numbers[0] * (numbers[7]*numbers[10] - numbers[6]*numbers[11])) - (numbers[4] * (numbers[3]*numbers[10] - numbers[2]*numbers[11])) + (numbers[8] * (numbers[3]*numbers[6] - numbers[2]*numbers[7])))
			dz = float64((numbers[0] * (numbers[5]*numbers[11] - numbers[7]*numbers[9])) - (numbers[4] * (numbers[1]*numbers[11] - numbers[3]*numbers[9])) + (numbers[8] * (numbers[1]*numbers[7] - numbers[3]*numbers[5])))

			// Prints either the solution, inconsistent, or dependent
			if d != 0 {
				fmt.Fprintln(w, "Solution:")
				fmt.Fprintf(w, "x = %.2f, y = %.2f, z = %.2f", dx/d, dy/d, dz/d)
			} else if d == 0 && dx == 0 && dy == 0 && dz == 0 {
				fmt.Fprintln(w, "Dependent - With Multiple Solutions")
			} else {
				fmt.Fprintln(w, "Inconsistent - No Solution")
			}
		}
	}
}

// IsInteger returns bool whether the string number can be converted to integer and also returns the converted integer
func IsInteger(num string) (bool, int) {
	if val, err := strconv.Atoi(num); err == nil {
		return true, val
	}
	return false, 0
}
