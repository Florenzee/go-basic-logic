package logic1

// function menerima parameter n, initial, dan step lalu mengembalikan variabel "slice" tipe slice 1D
func LogicAscDesc(n int, initial int, step int) (slice []int) {
	slice = make([]int, n)
	numPoint := initial

	for i := 0; i < n; i++ {
		slice[i] = numPoint

		if i <= n/2 {
			numPoint += step
		} else {
			numPoint -= step
		}
	}
	return slice
}
