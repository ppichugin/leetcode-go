package main

import "fmt"

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/807/
// LeetCode solution:
// Runtime: 65 ms, Memory: 6.9 MB

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := target - 1; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}
	return target == 0
}
