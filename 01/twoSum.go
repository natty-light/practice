package main

import (
	"fmt"
)

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0

	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var dict = make(map[int]int)
	for i, elem := range nums {

		if _, exists := dict[target-elem]; exists {
			return []int{i, dict[target-elem]}
		}
		dict[elem] = i
	}
	return []int{len(nums) - 1, dict[target-nums[len(nums)-1]]}
}
