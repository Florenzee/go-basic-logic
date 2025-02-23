package logic1

// n = jumlah index, init = nilai awal, step = nilai yg ditambahkan
func MatriksPattern(n int, init int, step int) []int {
	result := make([]int, n)
	value := init

	mid := n / 2 //cari index tengah

	// First half (incrementing)
	for i := 0; i < mid; i++ {
		result[i] = value
		value += step
	}
	// Middle index (continuing +2 if n is odd)
	result[mid] = value
	if n%2 == 1 {
		value += step
	}
	// Second half (decrementing)
	for i := mid; i < n; i++ {
		value -= step
		result[i] = value
	}
	return result
}
