package logic1

// function menerima parameter n, initial, dan step lalu mengembalikan variabel "slice" tipe slice 1D
func Logic1DescStep(n int, initial int, step int) (slice []int) {
	slice = make([]int, n)
	numPoint := initial
	for i := 0; i < n; i++ {
		slice[i] = numPoint
		numPoint -= step
	}
	return slice
}
