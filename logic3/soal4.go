package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks4(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			//baris genap diisi dari kanan ke kiri
			for j := n - 1; j >= 0; j-- {
				if i+j >= n-1 {
					result[i][j] = num
					num += 2
				}
			}
		} else {
			//baris ganjil diisi dari kiri ke kanan
			for j := 0; j < n; j++ {
				if i+j >= n-1 {
					result[i][j] = num
					num += 2
				}
			}
		}
	}
	return result
}
