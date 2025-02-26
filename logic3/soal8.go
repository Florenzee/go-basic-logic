package logic3

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
)

func Matriks8(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	nTengah := (n - 1) / 2
	num := 1

	// iterate only the left half columns, mirror the right half
	for col := 0; col < (n+1)/2; col++ {
		// iterate inside the top & bottom edges
		for row := nTengah - col; row <= nTengah+col; row++ {
			if col%2 == 1 { // if the col index is odd, fill from the top
				result[row][col] = num     // left half
				result[row][n-1-col] = num // mirror right half
			} else { // if the col index is even, fill from the bottom
				result[n-1-row][col] = num     // left half
				result[n-1-row][n-1-col] = num // mirror right half
			}
			num += 2

		}
	}
	return result
}
