package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var nums []int
	var input, median int

	_, err := fmt.Scan(&input)

	for err == nil {
		nums = append(nums, input)
		sort.Ints(nums)

		if len(nums)%2 == 1 {
			var index = int(math.Ceil(float64(len(nums) / 2)))
			median = nums[index]
		} else {
			median = int((nums[int(len(nums)/2)-1] + nums[int(len(nums)/2)]) / 2)
		}

		fmt.Println(median)

		_, err = fmt.Scan(&input)
	}
}
