package popcount

func PopCount(x uint64) int {
	var pc int
	for i := 0; i < 64; i++ {
		pc += int((x >> i) & 1)
	}
	return pc
}
