package logic2

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
)

func Matriks6(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][n-1-j] = start
				start += 2
			}
			//start += 3
		}
	}
	return result
}
