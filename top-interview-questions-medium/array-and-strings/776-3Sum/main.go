package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/

/*
Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

*/

// LeetCode solution took:
// Runtime: 62 ms
// Memory: 7.9 MB

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4} // test: [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	result := make([][]int, 0, l)

	for i := 0; i < l; i++ {
		// skipping duplicates
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			case sum < 0:
				j++
			default:
				k--
			}
		}
	}
	return result
}
