package popcount3

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var res byte
	for i := uint(0); i < 64; i++ {
		res += pc[byte(x>>(i))]
	}
	return int(res)
}
