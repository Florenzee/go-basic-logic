package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks5(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	mid := n / 2
	num := 29

	// Build the diamond shape
	for i := 0; i <= mid; i++ {
		for j := 0; j <= i; j++ {

			if i%2 == 0 {
				//baris ganjil diisi dari kanan ke kiri
				for j := n - 1; j >= 0; j-- {
					if i+j >= n-1 {
						result[i][j] = num
						num += 2
					}
				}
			}

			// Fill left and right sides symmetrically
			result[mid-j][mid-i] = num
			result[mid+j][mid-i] = num
			result[mid-j][mid+i] = num
			result[mid+j][mid+i] = num
			num -= 2
		}
	}

	return result
}
