package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/
// LeetCode Solution: Runtime 26 ms, Memory 7.3 MB

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"} // should be [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams(input))
}

func groupAnagrams(strs []string) [][]string {
	l := len(strs)
	m := make(map[string][]string, l)
	for i := range strs {
		str := strs[i]
		copyOfStr := []rune(str)
		sort.Slice(copyOfStr, func(i, j int) bool {
			return copyOfStr[j] < copyOfStr[i]
		})

		v := m[string(copyOfStr)]
		v = append(v, str)
		m[string(copyOfStr)] = v
	}

	v := make([][]string, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}

	return v
}
