package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks2(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	value := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			//baris genap isi dr kiri ke kanan
			for j := i; j < n; j++ {
				result[i][j] = value
				value += 2
			}
		} else {
			//baris ganjil isi dr kanan ke kiri
			for j := n - 1; j >= i; j-- {
				result[i][j] = value
				value += 2
			}
		}
	}
	return result
}
