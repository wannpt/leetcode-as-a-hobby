package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	smallest, small := math.MaxInt32, math.MaxInt32
	found := false
	for i := 0; i < len(nums); i++ {
		if nums[i] <= smallest {
			smallest = nums[i]
		} else if nums[i] <= small {
			small = nums[i]
		} else {
			found = true
		}

	}

	return found
}
func main() {
	ok := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	fmt.Println(ok)
}
