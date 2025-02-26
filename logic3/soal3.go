package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks3(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := -2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= n-1-i {
				if i == 0 {
					num += 3
					result[i][j] = num

				} else {
					if i%2 == 0 {
						num += 2
						result[i][j] = num

					} else {
						num += result[i][n-1-j-i]

					}
				}

			}

		}

	}
	return result
}
