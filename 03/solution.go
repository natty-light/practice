package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string = "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(in))
}

func lengthOfLongestSubstring(s string) int {
	var best, current int = 0, 0
	var dict = make(map[string]bool)
	for _, char := range strings.Split(s, "") {
		if _, exists := dict[char]; exists {
			if best > current {
				best = current
			}
			current = 1
			dict = make(map[string]bool)
			dict[char] = true
		} else {
			current += 1
			dict[char] = true
		}
	}
	if best > current {
		return best
	}
	return current
}
