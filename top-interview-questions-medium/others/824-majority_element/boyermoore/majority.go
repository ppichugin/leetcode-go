package boyermoore

// MajorityElement solution with Boyer-Moore Voting Algorithm.
// Time complexity : O(n). Space complexity: O(1)
func MajorityElement(nums []int) int {
	var count, candidate int

	for i, num := range nums {
		if i == 0 {
			candidate = num
		}
		count += counter(candidate, num)
	}
	return candidate
}

func counter(candidate, num int) int {
	if candidate == num {
		return 1
	}
	return -1
}
