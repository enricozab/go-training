package main

import "fmt"

func main() {
	var velocity, time int
	_, err := fmt.Scan(&velocity, &time)

	for err == nil {
		fmt.Println(2 * velocity * time)
		_, err = fmt.Scan(&velocity, &time)
	}
}
