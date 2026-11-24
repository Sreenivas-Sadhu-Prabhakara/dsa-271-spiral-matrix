package main

import "fmt"

// Spiral returns the matrix's values in clockwise spiral order: four
// boundary pointers shrink inward, with guards so a single leftover row or
// column isn't walked twice.
func Spiral(matrix [][]int) []int {
	out := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return out
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			out = append(out, matrix[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			out = append(out, matrix[r][right])
		}
		right--
		if top <= bottom { // guard: a single remaining row mustn't repeat
			for c := right; c >= left; c-- {
				out = append(out, matrix[bottom][c])
			}
			bottom--
		}
		if left <= right { // guard: a single remaining column mustn't repeat
			for r := bottom; r >= top; r-- {
				out = append(out, matrix[r][left])
			}
			left++
		}
	}
	return out
}

func main() {
	fmt.Println(Spiral([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
