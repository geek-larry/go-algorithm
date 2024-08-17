package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

func main() {
	array := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	target := 8
	a := searchMatrix(array, target)
	println(a)
}
