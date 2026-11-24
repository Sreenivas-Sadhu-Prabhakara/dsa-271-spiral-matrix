package main

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	cases := []struct {
		in   [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{7}}, []int{7}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},     // single row
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}}, // single column
		{[][]int{}, []int{}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
	}
	for _, c := range cases {
		if got := Spiral(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Spiral = %v, want %v", got, c.want)
		}
	}
}
