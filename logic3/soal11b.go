package logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Matriks11b(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	nTengah := (n - 1) / 2

	valueLeftMid := 1  //menyimpan sebagian nilai matriks kiri
	valueRightMid := 1 //menyimpan sebagian nilai matriks kanan

	for row := 0; row <= nTengah; row++ {
		//isi sebagian matriks kiri
		result[row][row] = valueLeftMid     // top half
		result[n-1-row][row] = valueLeftMid // mirror bottom half
		valueLeftMid += 2

		//isi sebagian matriks kanan
		for col := n - 1 - row; col < n; col++ {
			if row%2 == 1 {
				//row ganjil isi dari kiri
				result[row][col] = valueRightMid                 // top half left quarter
				result[n-1-row][col] = valueRightMid             // mirror bottom half left quarter
				result[row][(n-1-col)+(n-1)] = valueRightMid     // mirror top half right quarter
				result[n-1-row][(n-1-col)+(n-1)] = valueRightMid // mirror bottom half right quarter
			} else {
				//row genap isi dari kanan
				result[row][(n-1-col)+(n-1-row)] = valueRightMid     // top half left quarter
				result[n-1-row][(n-1-col)+(n-1-row)] = valueRightMid // mirror bottom half left quarter
				result[row][col+row] = valueRightMid                 // mirror top half right quarter
				result[n-1-row][col+row] = valueRightMid             // mirror bottom half right quarter
			}
			valueRightMid += 2
		}
	}
	return result
}
