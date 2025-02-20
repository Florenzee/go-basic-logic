package logic2

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks8(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				result[i][j] = 2*j + 1
			}
		}
	}
	return result

}
