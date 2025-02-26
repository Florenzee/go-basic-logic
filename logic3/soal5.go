package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks5(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	mid := n / 2
	values := 29

	// Build the diamond shape
	for i := 0; i <= mid; i++ {
		for j := 0; j <= i; j++ {

			//

			// Fill left and right sides symmetrically
			result[mid-j][mid-i] = values
			result[mid+j][mid-i] = values
			result[mid-j][mid+i] = values
			result[mid+j][mid+i] = values
			values -= 2
		}
	}

	return result
}
