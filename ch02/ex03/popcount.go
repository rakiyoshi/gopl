package popcount

// pc[i] is population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var popcount int
	for i := 0; i < 8; i++ {
		popcount += int(pc[byte(x>>(i*8))])
	}
	return popcount
}
