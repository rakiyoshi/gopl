package popcount

func PopCount(x uint64) int {
	var pc int
	for x != 0 {
		x = x & (x - 1)
		pc++
	}
	return pc
}
