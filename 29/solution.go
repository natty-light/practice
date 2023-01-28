package main

import (
	"fmt"
	"math"
)

func main() {
	var dividend, divisor = 1, -1
	fmt.Println(divide(dividend, divisor))
	dividend, divisor = -1, 1
	fmt.Println(divide(dividend, divisor))
	dividend = -2147483648
	divisor = -1
	fmt.Println(divide(dividend, divisor))
	dividend = 2147483647
	divisor = 1
	fmt.Println(divide(dividend, divisor))
}

// recursion failed me
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 1 || divisor == -1 {
		return dividend * divisor
	}
	if (dividend < 0) != (divisor < 0) {
		if dividend+divisor < 0 {
			return 0
		} else {
			return -1 + divide(dividend+divisor, divisor)
		}
	} else {
		if dividend-divisor < 0 {
			return 0
		} else {
			return 1 + divide(dividend-divisor, divisor)
		}
	}
}
