package logic3

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
)

func Matriks8(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := 1
	nTengah := (n - 1) / 2

	for i := 0; i <= nTengah; i++ {
		for j := 0; j < n; j++ {
			if i+j >= nTengah && j-i <= nTengah {
				if i%2 == 1 {
					result[j][i] = num
					result[j][n-1-i] = num
				} else {
					result[n-1-j][i] = num
					result[n-1-j][n-1-i] = num
				}
			}
		}
	}
	return result
}
