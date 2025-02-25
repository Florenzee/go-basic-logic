package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks3(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	num := 49

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < n-i; j++ {
				result[i][j] = num
				num += 3
			}
		} else {
			for j := n - 1 - i; j >= 0; j-- {
				result[i][j] = num
				num += 2
			}
		}
	}

	//for j := 0; j <= n-1; j++ {
	//	if i+j >= n-1 {
	//		//genap isi dari kiri ke kanan
	//		if i%2 == 0 {
	//			result[j][i] = num
	//			num += 3
	//			//ganjil isi dari kanan ke kiri
	//		} else {
	//			result[j][i] = num
	//			num += 2
	//		}
	//	}
	//}
	return result
}
