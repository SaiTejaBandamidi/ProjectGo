package main

import (
	"fmt"
)

func pairSum(nums []int, target int) []int {
	numMap := make(map[int]int) // HashMap to store numbers and their indices

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i // Store the current number with its index
	}
	return nil // Should never reach here as per problem constraints
}

func main() {
	// Test cases
	fmt.Println(pairSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(pairSum([]int{3, 2, 4}, 6))

	// Output: [1, 2]
	fmt.Println(pairSum([]int{5, 8, 6, 1}, 7)) //	Output: [2,3]
}
