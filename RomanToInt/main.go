package main

// https://leetcode.com/problems/roman-to-integer/

import "fmt"

func main() {
	romanNumeral := "MCMXCIV"
	integerValue := romanToInt(romanNumeral)
	fmt.Println(integerValue)
}

func getSymbol(char byte) int {
	switch char {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currValue := getSymbol(s[i])
		if currValue < prevValue {
			result -= currValue
		} else {
			result += currValue
		}
		prevValue = currValue
	}
	return result
}
