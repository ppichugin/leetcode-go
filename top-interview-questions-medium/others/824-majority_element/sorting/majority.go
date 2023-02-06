package sorting

import "sort"

// MajorityElement solution based on sorting algorithm.
// Time complexity: O(n * lgn). Space complexity: O(1)
func MajorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
