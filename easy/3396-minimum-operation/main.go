package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumOperation([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
	fmt.Println(minimumOperation([]int{4, 5, 6, 4, 4}))
	fmt.Println(minimumOperation([]int{6, 7, 8, 9}))
	fmt.Println(minimumOperation([]int{10, 12, 12, 6, 6}))
	fmt.Println(minimumOperation([]int{4, 5, 6, 4, 4}))
	fmt.Println(minimumOperation([]int{2, 2, 6, 13, 6, 13}))
	fmt.Println(minimumOperation([]int{6, 6, 13, 2, 6, 13}))
}

// Minimum Number of Operations to Make Elements in Array Distinct
func minimumOperation(nums []int) int {
	operations := 0

	for {
		// Create a map to count the frequency of each number
		count := make(map[int]int)
		for _, num := range nums {
			count[num]++
		}
		fmt.Println("Count:", count)

		// Check if all elements are already distinct
		if len(count) == len(nums) {
			break
		}

		// Handle arrays with 3+ elements
		if len(nums) >= 3 {
			// Find at least one duplicate and remove first 3 elements
			duplicateFound := false
			for _, freq := range count {
				if freq > 1 {
					duplicateFound = true
					break
				}
			}

			if duplicateFound {
				nums = slices.Delete(nums, 0, 3)
				operations++
				continue // Restart the loop with the updated array
			}
		} else {
			// For arrays with fewer than 3 elements
			for _, freq := range count {
				if freq > 1 {
					operations++
					break
				}
			}
			break
		}
	}

	return operations
}
