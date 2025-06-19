package main

import "fmt"

func getUniquePaths(rows int, cols int) int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
				continue
			}

			above := matrix[i-1][j]
			left := matrix[i][j-1]

			matrix[i][j] = left + above
		}
	}

	return matrix[rows-1][cols-1]
}

func main() {
	fmt.Println("Expected: 3, Returned: ", getUniquePaths(2, 3))
	fmt.Println("Expected: 10, Returned: ", getUniquePaths(3, 4))
	fmt.Println("Expected: 210, Returned: ", getUniquePaths(5, 7))
}
