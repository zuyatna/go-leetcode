package main

// https://leetcode.com/problems/longest-common-prefix/description/

import "fmt"

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"ab", "a"})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	fmt.Println(prefix)
	for i := 1; i < len(strs); i++ {
		fmt.Println(i, " < ", len(strs))
		for len(prefix) > 0 {
			fmt.Println(len(prefix), " > 0")
			fmt.Println(strs[i], " ", prefix)
			fmt.Println(len(strs[i]), " >= ", len(prefix), " && ", strs[i], " len[:", len(prefix), "] == ", prefix)
			if len(strs[i]) >= len(prefix) && strs[i][:len(prefix)] == prefix {
				fmt.Println(len(strs[i]), " >= ", len(prefix), " && ", strs[i][:len(prefix)], " == ", prefix)
				break
			}
			prefix = prefix[:len(prefix)-1]
			fmt.Println(prefix, " -1")
		}
	}

	fmt.Println(prefix)
	return prefix
}
