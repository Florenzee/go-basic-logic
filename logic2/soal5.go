package logic2

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks5(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	value := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				result[i][j] = value
				value += 2
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				result[i][j] = value
				value += 2
			}
		}
	}
	return result
}
