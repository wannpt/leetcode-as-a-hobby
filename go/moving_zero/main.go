package main

import "fmt"

func movingZeroes(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		fmt.Printf("test %v \n", &nums[i])
	}
	return nums
}

func main() {

	tmp := movingZeroes([]int{0, 1, 0, 3, 12})
	fmt.Printf("%v", tmp)
}
