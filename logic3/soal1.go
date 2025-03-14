package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks1(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	value := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			//isi dari kiri
			for j := 0; j <= i; j++ {
				result[i][j] = value
				value += 2
			}
		} else {
			//isi dari kanan
			for j := i; j >= 0; j-- {
				result[i][j] = value
				value += 2
			}
		}
	}
	return result
}
