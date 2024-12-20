package main

// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description/?envType=daily-question&envId=2024-11-14

import "fmt"

func main() {
	n := 6
	quantities := []int{11, 6}
	result := minimizedMaximum1(n, quantities)
	fmt.Println(result)
}

func minimizedMaximum1(n int, quantities []int) int {
	low, high := 1, 0
	// find initial high
	for _, q := range quantities {
		// if the current quantity (q) is greater than high, updates high to q. This ensures high is large enough to hold any single item.
		if q > high {
			high = q
		}
	}

	// binary search
	for low < high {
		mid := (low + high) / 2 // calculates the middle capacity to check
		needed := 0             // store the total number of containers required if all items were assigned a capacity of mid
		for _, q := range quantities {
			// calculates the number of containers needed for the current item (q) using the current capacity (mid)
			// the [(q + mid - 1) / mid] expression performs integer division,
			// effectively calculating the ceiling (minimum number of containers) required to hold q items with a capacity of mid.
			needed += (q + mid - 1) / mid
		}
		// update search space based on needed
		if needed <= n {
			high = mid
		} else {
			low = mid + 1 // update to consider higher capacities that might fit within n containers
		}
	}

	return low // low will hold the minimized maximum capacity
}
