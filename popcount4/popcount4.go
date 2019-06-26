package popcount4

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var res int
	for x != 0 {
		res += 1
		x = x & (x - 1)
	}
	return res
}
