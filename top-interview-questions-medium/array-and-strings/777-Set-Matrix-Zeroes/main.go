package main

import "fmt"

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/
// LeetCode solution: Runtime: 12 ms, Memory 6.3 MB

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	isFirstCol := false

	// Scanning rows
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			isFirstCol = true
		}

		// Scanning columns,
		// marking first cell of each column
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Filling with zeros based on cells info
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Taking into consideration the first cell of matrix
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Filling first column with zero
	if isFirstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
