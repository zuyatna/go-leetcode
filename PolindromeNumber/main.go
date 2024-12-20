package main

// https://leetcode.com/problems/palindrome-number/description/
// Given an integer x, return true if x is a palindrome, and false otherwise.

// An integer is a palindrome when it reads the same forward and backward.

import (
	"fmt"
	"strconv"
)

func main() {
	x := -121
	if isPolindrome(x) {
		fmt.Println("the number is polindrome")
	} else {
		fmt.Println("the number is not polindrome")
	}
}

func isPolindrome(x int) bool {
	// first, I think I need to convert integer to string, after that I will slice the string and compare that string
	str := strconv.Itoa(x)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}
