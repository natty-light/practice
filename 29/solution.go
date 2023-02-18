package main

import (
	"fmt"
	"math"
)

func main() {
	var dividend, divisor = 10, 3
	var result = divide(dividend, divisor)
	fmt.Printf("Expected: 2147483647 \n")
	fmt.Printf("%d / %d = %d \n\n", dividend, divisor, result)
}

func divide(dividend int, divisor int) int {
	var sign, result, accumulator int

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Compute and store the sign
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	} else {
		sign = -1
	}

	// Get the easy one out of the way first
	if dividend == divisor || -dividend == divisor {
		return 1 * sign
	}
	if dividend == 0 {
		return 0
	}

	// Peel off sign
	dividend, divisor = abs(dividend), abs(divisor)
	result = 0
	// While the dividend is bigger than the divisor
	for dividend-divisor >= 0 {
		// start the accumulator
		accumulator = 0
		// Bit shift by 1 + accumulator. Equivalent to multiplying by 2
		for dividend-(divisor<<1<<accumulator) >= 0 {
			// once accumulator * 2 > divisor, break out
			accumulator++
		}
		//keep track of the result by adding 2*accumulator
		result += 1 << accumulator
		// Remove lower the divisor by the divisor 2*accumulator, reset the accumulator so as to only divide by 1*divisor on line 78
		dividend -= divisor << accumulator
	}
	return result * sign
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
