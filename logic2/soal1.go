package logic2

// function matriks1 menerima parameter n dan mengembalikan variabel result tipe slice 2D
func Matriks1(n int) (result [][]int) {
	//membuat slice n*n
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		//mengisi setiap matriks dengan nilai sesuai soal
		for j := 0; j < n; j++ {
			result[i][j] = 1 + 2*j
		}
	}

	////mengisi slice
	//num := 1
	//for row := 0; row < n; row++ {
	//	for col := 0; col < n; col++ {
	//		//proses mengisi
	//		result[row][col] = num
	//	}
	//}

	return result //mengembalikan result
}
