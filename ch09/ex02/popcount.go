package popcount

import "sync"

// pc[i] is population count of i
var pc [256]byte

func initPc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var (
	loadInitPcOnce sync.Once
)

func PopCount(x uint64) int {
	loadInitPcOnce.Do(initPc)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
