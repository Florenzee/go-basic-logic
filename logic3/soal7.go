package logic3

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
)

func Matriks7(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := 1
	nTengah := (n - 1) / 2

	for i := 0; i <= nTengah; i++ {
		for j := 0; j < n; j++ {
			if i+j >= nTengah && j-i <= nTengah {
				//result[j][i] = num

				if j%2 == 1 {
					result[i][j] = num
					result[n-1-i][j] = num
				} else {
					result[i][n-1-j] = num
					result[n-1-i][n-1-j] = num
				}
				num += 2
			}
		}
	}
	return result
}
