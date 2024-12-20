package main

// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description/?envType=daily-question&envId=2024-11-14

import "fmt"

func main() {
	n := 7
	quantities := []int{15, 10, 10}
	result := minimizedMaximum2(n, quantities)
	fmt.Println(result)
}

func minimizedMaximum2(n int, quantities []int) int {
	canDistribute := func(x int) bool {
		//pointer to the first not fully distributed product type
		j := 0
		// remaining quantity of the j-th product type
		remaining := quantities[j]

		// loop through each store
		for i := 0; i < n; i++ {
			// check if the remaining quantity of the j-th product type
			// can be fully distributed to the  i-th store
			if remaining <= x {
				// if yes, move the pointer to the next product type
				j++
				// check if all products have been distributed
				if j == len(quantities) {
					return true
				} else {
					remaining = quantities[j]
				}
			} else {
				// distributed the maximum possible quantity (x) to the i-th store
				remaining -= x
			}
		}
		return false
	}

	// initialize the boundaries of the binary search
	left, right := 0, 0

	// find the maximum element in quantities
	for _, quantity := range quantities {
		if quantity > right {
			right = quantity
		}
	}

	// perform binary search until the boundaries coverage
	for left < right {
		middle := (left + right) / 2

		if canDistribute(middle) {
			// try smaller maximum
			right = middle
		} else {
			// increase the minimum possible maximum
			left = middle + 1
		}
	}

	return left
}
