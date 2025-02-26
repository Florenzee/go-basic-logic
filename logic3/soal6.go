package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks6(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (j >= i && i+j+1 <= n) || (j <= i && i+j+1 >= n) {
				result[i][j] = num
			}
		}
	}

	return result
}
