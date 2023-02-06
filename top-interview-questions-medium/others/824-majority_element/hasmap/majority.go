package hasmap

// MajorityElement solution based inb hashmap.
// Time complexity : O(n). Space complexity: O(n)
func MajorityElement(nums []int) int {
	var maxKey, maxValue int
	m := countNums(nums)

	for key, value := range m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}

func countNums(nums []int) map[int]int {
	m := make(map[int]int, len(nums))

	for _, num := range nums {
		m[num]++
	}

	return m
}
