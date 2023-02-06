package main

import (
	"fmt"

	"leetcode-go/top-interview-questions-medium/others/824-majority_element/boyermoore"
	"leetcode-go/top-interview-questions-medium/others/824-majority_element/hasmap"
	"leetcode-go/top-interview-questions-medium/others/824-majority_element/sorting"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/824/

func main() {
	fmt.Println(boyermoore.MajorityElement([]int{3, 2, 3}))
	fmt.Println(hasmap.MajorityElement([]int{3, 2, 3}))
	fmt.Println(sorting.MajorityElement([]int{3, 2, 3}))
}
