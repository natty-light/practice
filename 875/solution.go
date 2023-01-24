package main

import (
	"fmt"
	"math"
)

func main() {
	var piles = []int{3, 6, 7, 11}
	var h int = 8

	var k int = minEatingSpeed(piles, h)

	fmt.Printf("k = %d\n", k)

}

func minEatingSpeed(piles []int, h int) (k int) {
	var left, right int = 1, max(piles)
	for left < right {
		mid := left + (right-left)/2
		fmt.Printf("left = %d, right = %d, mid = %d \n", left, right, mid)
		if canEat(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canEat(piles []int, k int, h int) bool {
	var elapsedHours = 0
	for _, pile := range piles {
		elapsedHours += int(math.Ceil(float64(pile) / float64(k)))
		if elapsedHours > h {
			return false
		}
	}
	return true
}

func max(slice []int) (res int) {
	res = -1
	for _, elem := range slice {
		if elem > res {
			res = elem
		}
	}
	return res
}
