package logic2

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks4(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	value := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			result[i][j] = value
			if j == 0 || i == 0 || j == n-1 {
				value += 3
			} else {
				value += 2
			}
		}
	}
	return result
}
